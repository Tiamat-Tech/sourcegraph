package conf

import (
	"encoding/json"
	"fmt"
	"strings"

	"github.com/sourcegraph/sourcegraph/schema"
)

func validateCustomRaw(normalizedInput []byte) (validationErrors []string, err error) {
	var cfg schema.SiteConfiguration
	if err := json.Unmarshal(normalizedInput, &cfg); err != nil {
		return nil, err
	}
	return validateCustom(cfg)
}

// validateCustom validates the site config using custom validation steps that are not
// able to be expressed in the JSON Schema.
func validateCustom(cfg schema.SiteConfiguration) (validationErrors []string, err error) {

	invalid := func(msg string) {
		validationErrors = append(validationErrors, msg)
	}

	if cfg.AuthProvider != "" && len(cfg.AuthProviders) > 0 {
		invalid(`auth.providers takes precedence over auth.provider (deprecated) when both are set (auth.provider is IGNORED in that case)`)
	} else if cfg.AuthProvider != "" {
		invalid(`auth.provider is deprecated; use auth.providers instead`)
	}
	if len(cfg.AuthProviders) >= 2 {
		invalid(`auth.providers supports only a single entry (entries other than the first are IGNORED)`)
	}

	authProvider := authProvider(&cfg)
	if authProvider == (schema.AuthProviders{}) {
		invalid("no auth provider set (all access will be forbidden)")
	}
	if cfg.AuthAllowSignup && authProvider.Builtin == nil {
		invalid(fmt.Sprintf("auth.allowSignup requires auth provider \"builtin\" (got %q)", cfg.AuthProvider))
	}
	if cfg.AuthAllowSignup {
		invalid(fmt.Sprintf(`auth.allowSignup is deprecated; use "auth.providers" with an entry of {"type":"builtin","allowSignup":true} instead`))
	}

	if (authProvider.Openidconnect != nil || authProvider.Saml != nil) && cfg.AppURL == "" {
		invalid(`auth providers "openidconnect" and "saml" require appURL to be set to the external URL of your site (example: https://sourcegraph.example.com)`)
	}

	{
		hasOldOIDC := cfg.OidcProvider != "" || cfg.OidcClientID != "" || cfg.OidcClientSecret != "" || cfg.OidcEmailDomain != ""
		hasSingularOIDC := cfg.AuthOpenIDConnect != nil
		if hasOldOIDC {
			invalid(`oidc* properties are deprecated; use auth provider "openidconnect" instead`)
		}
		if cfg.AuthProvider == "openidconnect" && !hasSingularOIDC {
			invalid(`auth.openIDConnect must be configured when auth.provider == "openidconnect"`)
		}
		if hasOldOIDC && cfg.AuthProvider != "openidconnect" {
			invalid(`must set auth.provider == "openidconnect" for oidc* config to take effect (also, oidc* config is deprecated; see other message to that effect)`)
		}
		if hasSingularOIDC && cfg.AuthProvider != "openidconnect" {
			invalid(`must set auth.provider == "openidconnect" for auth.openIDConnect config to take effect`)
		}
	}

	if authProvider.Openidconnect != nil && authProvider.Openidconnect.OverrideToken != "" {
		invalid(`OpenID Connect auth provider "overrideToken" is deprecated (because it applies to all auth providers, not just OIDC); use OVERRIDE_AUTH_SECRET env var instead`)
	}

	{
		hasOldSAML := cfg.SamlIDProviderMetadataURL != "" || cfg.SamlSPCert != "" || cfg.SamlSPKey != ""
		hasSingularSAML := cfg.AuthSaml != nil
		if hasOldSAML {
			invalid(`saml* properties are deprecated; use auth provider "saml" instead`)
		}
		if cfg.AuthProvider == "saml" && !hasSingularSAML {
			invalid(`auth.saml must be configured when auth.provider == "saml"`)
		}
		if hasOldSAML && cfg.AuthProvider != "saml" {
			invalid(`must set auth.provider == "saml" for saml* config to take effect (also, saml* config is deprecated; see other message to that effect)`)
		}
		if hasSingularSAML && cfg.AuthProvider != "saml" {
			invalid(`must set auth.provider == "saml" for auth.saml config to take effect`)
		}
	}

	{
		hasSingularAuthHTTPHeader := cfg.AuthUserIdentityHTTPHeader != ""
		if cfg.AuthProvider == "http-header" && !hasSingularAuthHTTPHeader {
			invalid(`auth.userIdentityHTTPHeader must be configured when auth.provider == "http-header"`)
		}
		if hasSingularAuthHTTPHeader && cfg.AuthProvider != "http-header" {
			invalid(`must set auth.provider == "http-header" for auth.userIdentityHTTPHeader config to take effect`)
		}
	}

	{
		hasSMTP := cfg.EmailSmtp != nil
		hasSMTPAuth := cfg.EmailSmtp != nil && cfg.EmailSmtp.Authentication != "none"
		if hasSMTP && cfg.EmailAddress == "" {
			invalid(`should set email.address because email.smtp is set`)
		}
		if hasSMTPAuth && (cfg.EmailSmtp.Username == "" && cfg.EmailSmtp.Password == "") {
			invalid(`must set email.smtp username and password for email.smtp authentication`)
		}
	}

	{
		for _, phabCfg := range cfg.Phabricator {
			if len(phabCfg.Repos) == 0 && phabCfg.Token == "" {
				invalid(`each phabricator instance must have either "token" or "repos" set`)
			}
		}
	}

	for _, bbsCfg := range cfg.BitbucketServer {
		if bbsCfg.Token != "" && (bbsCfg.Username != "" || bbsCfg.Password != "") {
			invalid("for Bitbucket Server, specify either a token or a username/password, not both")
		} else if bbsCfg.Token == "" && bbsCfg.Username == "" && bbsCfg.Password == "" {
			invalid("for Bitbucket Server, you must specify either a token or a username/password to authenticate")
		}
	}

	for _, ghCfg := range cfg.Github {
		if ghCfg.PreemptivelyClone {
			invalid(`github[].preemptivelyClone is deprecated; use initialRepositoryEnablement instead`)
		}
	}

	for _, c := range cfg.Gitlab {
		if strings.Contains(c.Url, "example.com") {
			invalid(fmt.Sprintf(`invalid GitLab URL detected: %s (did you forget to remove "example.com"?)`, c.Url))
		}
	}

	return validationErrors, nil
}
