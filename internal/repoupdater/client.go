package repoupdater

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/opentracing-contrib/go-stdlib/nethttp"
	"github.com/opentracing/opentracing-go/ext"
	"github.com/pkg/errors"

	"github.com/sourcegraph/sourcegraph/internal/api"
	"github.com/sourcegraph/sourcegraph/internal/env"
	"github.com/sourcegraph/sourcegraph/internal/metrics"
	"github.com/sourcegraph/sourcegraph/internal/repoupdater/protocol"
	"github.com/sourcegraph/sourcegraph/internal/trace/ot"
)

var requestMeter = metrics.NewRequestMeter("repoupdater", "Total number of requests sent to repoupdater.")

// DefaultClient is the default Client. Unless overwritten, it is
// connected to the server specified by the REPO_UPDATER_URL
// environment variable.
var DefaultClient = NewClient(env.Get("REPO_UPDATER_URL", "http://repo-updater:3182", "repo-updater server URL"))

// client is a repoupdater client.
type client struct {
	// URL to repoupdater server.
	URL string

	// HTTP client to use
	HTTPClient *http.Client
}

func NewClient(serverURL string) *client {
	return &client{
		URL: serverURL,
		HTTPClient: &http.Client{
			// ot.Transport will propagate opentracing spans and whether or not to trace
			Transport: &ot.Transport{
				RoundTripper: requestMeter.Transport(&http.Transport{
					// Default is 2, but we can send many concurrent requests
					MaxIdleConnsPerHost: 500,
				}, func(u *url.URL) string {
					// break it down by API function call (ie "/repo-update-scheduler-info", "/repo-lookup", etc)
					return u.Path
				}),
			},
		},
	}
}

// RepoUpdateSchedulerInfo returns information about the state of the repo in the update scheduler.
func (c *client) RepoUpdateSchedulerInfo(ctx context.Context, args protocol.RepoUpdateSchedulerInfoArgs) (result *protocol.RepoUpdateSchedulerInfoResult, err error) {
	resp, err := c.httpPost(ctx, "repo-update-scheduler-info", args)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		stack := fmt.Sprintf("RepoScheduleInfo: %+v", args)
		return nil, errors.Wrap(fmt.Errorf("http status %d", resp.StatusCode), stack)
	}
	defer resp.Body.Close()
	err = json.NewDecoder(resp.Body).Decode(&result)
	return result, err
}

// MockRepoLookup mocks (*client).RepoLookup for tests.
var MockRepoLookup func(protocol.RepoLookupArgs) (*protocol.RepoLookupResult, error)

// RepoLookup retrieves information about the repository on repoupdater.
func (c *client) RepoLookup(ctx context.Context, args protocol.RepoLookupArgs) (result *protocol.RepoLookupResult, err error) {
	if MockRepoLookup != nil {
		return MockRepoLookup(args)
	}

	span, ctx := ot.StartSpanFromContext(ctx, "client.RepoLookup")
	defer func() {
		if result != nil {
			span.SetTag("found", result.Repo != nil)
		}
		if err != nil {
			ext.Error.Set(span, true)
			span.SetTag("err", err.Error())
		}
		span.Finish()
	}()
	if args.Repo != "" {
		span.SetTag("Repo", string(args.Repo))
	}

	resp, err := c.httpPost(ctx, "repo-lookup", args)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		// best-effort inclusion of body in error message
		body, _ := ioutil.ReadAll(io.LimitReader(resp.Body, 200))
		return nil, errors.Errorf("RepoLookup for %+v failed with http status %d: %s", args, resp.StatusCode, string(body))
	}

	err = json.NewDecoder(resp.Body).Decode(&result)
	if err == nil && result != nil {
		switch {
		case result.ErrorNotFound:
			err = &ErrNotFound{
				Repo:       args.Repo,
				IsNotFound: true,
			}
		case result.ErrorUnauthorized:
			err = &ErrUnauthorized{
				Repo:    args.Repo,
				NoAuthz: true,
			}
		case result.ErrorTemporarilyUnavailable:
			err = &ErrTemporary{
				Repo:        args.Repo,
				IsTemporary: true,
			}
		}
	}
	return result, err
}

// MockEnqueueRepoUpdate mocks (*client).EnqueueRepoUpdate for tests.
var MockEnqueueRepoUpdate func(ctx context.Context, repo api.RepoName) (*protocol.RepoUpdateResponse, error)

// EnqueueRepoUpdate requests that the named repository be updated in the near
// future. It does not wait for the update.
func (c *client) EnqueueRepoUpdate(ctx context.Context, repo api.RepoName) (*protocol.RepoUpdateResponse, error) {
	if MockEnqueueRepoUpdate != nil {
		return MockEnqueueRepoUpdate(ctx, repo)
	}

	req := &protocol.RepoUpdateRequest{
		Repo: repo,
	}

	resp, err := c.httpPost(ctx, "enqueue-repo-update", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var res protocol.RepoUpdateResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

// MockEnqueueChangesetSync mocks (*client).EnqueueChangesetSync for tests.
var MockEnqueueChangesetSync func(ctx context.Context, ids []int64) error

func (c *client) EnqueueChangesetSync(ctx context.Context, ids []int64) error {
	if MockEnqueueChangesetSync != nil {
		return MockEnqueueChangesetSync(ctx, ids)
	}

	req := protocol.ChangesetSyncRequest{IDs: ids}
	resp, err := c.httpPost(ctx, "enqueue-changeset-sync", req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "failed to read response body")
	}

	var res protocol.ChangesetSyncResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return err
	}

	if res.Error == "" {
		return nil
	}
	return errors.New(res.Error)
}

func (c *client) SchedulePermsSync(ctx context.Context, args protocol.PermsSyncRequest) error {
	resp, err := c.httpPost(ctx, "schedule-perms-sync", args)
	if err != nil {
		return err
	}
	defer func() { _ = resp.Body.Close() }()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "read response body")
	}

	var res protocol.PermsSyncResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return err
	}

	if res.Error == "" {
		return nil
	}
	return errors.New(res.Error)
}

// SyncExternalService requests the given external service to be synced.
func (c *client) SyncExternalService(ctx context.Context, svc api.ExternalService) (*protocol.ExternalServiceSyncResult, error) {
	req := &protocol.ExternalServiceSyncRequest{ExternalService: svc}
	resp, err := c.httpPost(ctx, "sync-external-service", req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var result protocol.ExternalServiceSyncResult
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		// TODO(tsenart): Use response type for unmarshalling errors too.
		// This needs to be done after rolling out the response type in prod.
		return nil, errors.New(string(bs))
	} else if len(bs) == 0 {
		// TODO(keegancsmith): Remove once repo-updater update is rolled out.
		result.ExternalService = svc
		return &result, nil
	} else if err = json.Unmarshal(bs, &result); err != nil {
		return nil, err
	}

	if result.Error != "" {
		return nil, errors.New(result.Error)
	}
	return &result, nil
}

// RepoExternalServices requests the external services associated with a
// repository with the given id.
func (c *client) RepoExternalServices(ctx context.Context, id api.RepoID) ([]api.ExternalService, error) {
	req := protocol.RepoExternalServicesRequest{ID: id}
	resp, err := c.httpPost(ctx, "repo-external-services", &req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var res protocol.RepoExternalServicesResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return nil, err
	}

	return res.ExternalServices, nil
}

// ExcludeRepo adds the repository with the given id to all of the
// external services exclude lists that match its kind.
func (c *client) ExcludeRepo(ctx context.Context, id api.RepoID) (*protocol.ExcludeRepoResponse, error) {
	if id == 0 {
		return &protocol.ExcludeRepoResponse{}, nil
	}

	req := protocol.ExcludeRepoRequest{ID: id}
	resp, err := c.httpPost(ctx, "exclude-repo", &req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "failed to read response body")
	}

	var res protocol.ExcludeRepoResponse
	if resp.StatusCode < 200 || resp.StatusCode >= 400 {
		return nil, errors.New(string(bs))
	} else if err = json.Unmarshal(bs, &res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (c *client) httpPost(ctx context.Context, method string, payload interface{}) (resp *http.Response, err error) {
	reqBody, err := json.Marshal(payload)
	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", c.URL+"/"+method, bytes.NewReader(reqBody))
	if err != nil {
		return nil, err
	}

	return c.do(ctx, req)
}

func (c *client) do(ctx context.Context, req *http.Request) (_ *http.Response, err error) {
	span, ctx := ot.StartSpanFromContext(ctx, "client.do")
	defer func() {
		if err != nil {
			ext.Error.Set(span, true)
			span.SetTag("err", err.Error())
		}
		span.Finish()
	}()

	req.Header.Set("Content-Type", "application/json")

	req = req.WithContext(ctx)
	req, ht := nethttp.TraceRequest(span.Tracer(), req,
		nethttp.OperationName("RepoUpdater Client"),
		nethttp.ClientTrace(false))
	defer ht.Finish()

	if c.HTTPClient != nil {
		return c.HTTPClient.Do(req)
	}
	return http.DefaultClient.Do(req)
}
