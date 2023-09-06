// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package ratelimit

import (
	"context"
	"sync"
	"time"

	redispool "github.com/sourcegraph/sourcegraph/internal/redispool"
)

// MockRateLimiter is a mock implementation of the RateLimiter interface
// (from the package github.com/sourcegraph/sourcegraph/internal/redispool)
// used for unit testing.
type MockRateLimiter struct {
	// GetTokenFunc is an instance of a mock function object controlling the
	// behavior of the method GetToken.
	GetTokenFunc *RateLimiterGetTokenFunc
	// SetTokenBucketConfigFunc is an instance of a mock function object
	// controlling the behavior of the method SetTokenBucketConfig.
	SetTokenBucketConfigFunc *RateLimiterSetTokenBucketConfigFunc
}

// NewMockRateLimiter creates a new mock of the RateLimiter interface. All
// methods return zero values for all results, unless overwritten.
func NewMockRateLimiter() *MockRateLimiter {
	return &MockRateLimiter{
		GetTokenFunc: &RateLimiterGetTokenFunc{
			defaultHook: func(context.Context, string) (r0 error) {
				return
			},
		},
		SetTokenBucketConfigFunc: &RateLimiterSetTokenBucketConfigFunc{
			defaultHook: func(context.Context, string, int32, time.Duration) (r0 error) {
				return
			},
		},
	}
}

// NewStrictMockRateLimiter creates a new mock of the RateLimiter interface.
// All methods panic on invocation, unless overwritten.
func NewStrictMockRateLimiter() *MockRateLimiter {
	return &MockRateLimiter{
		GetTokenFunc: &RateLimiterGetTokenFunc{
			defaultHook: func(context.Context, string) error {
				panic("unexpected invocation of MockRateLimiter.GetToken")
			},
		},
		SetTokenBucketConfigFunc: &RateLimiterSetTokenBucketConfigFunc{
			defaultHook: func(context.Context, string, int32, time.Duration) error {
				panic("unexpected invocation of MockRateLimiter.SetTokenBucketConfig")
			},
		},
	}
}

// NewMockRateLimiterFrom creates a new mock of the MockRateLimiter
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockRateLimiterFrom(i redispool.RateLimiter) *MockRateLimiter {
	return &MockRateLimiter{
		GetTokenFunc: &RateLimiterGetTokenFunc{
			defaultHook: i.GetToken,
		},
		SetTokenBucketConfigFunc: &RateLimiterSetTokenBucketConfigFunc{
			defaultHook: i.SetTokenBucketConfig,
		},
	}
}

// RateLimiterGetTokenFunc describes the behavior when the GetToken method
// of the parent MockRateLimiter instance is invoked.
type RateLimiterGetTokenFunc struct {
	defaultHook func(context.Context, string) error
	hooks       []func(context.Context, string) error
	history     []RateLimiterGetTokenFuncCall
	mutex       sync.Mutex
}

// GetToken delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockRateLimiter) GetToken(v0 context.Context, v1 string) error {
	r0 := m.GetTokenFunc.nextHook()(v0, v1)
	m.GetTokenFunc.appendCall(RateLimiterGetTokenFuncCall{v0, v1, r0})
	return r0
}

// SetDefaultHook sets function that is called when the GetToken method of
// the parent MockRateLimiter instance is invoked and the hook queue is
// empty.
func (f *RateLimiterGetTokenFunc) SetDefaultHook(hook func(context.Context, string) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetToken method of the parent MockRateLimiter instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *RateLimiterGetTokenFunc) PushHook(hook func(context.Context, string) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RateLimiterGetTokenFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, string) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RateLimiterGetTokenFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, string) error {
		return r0
	})
}

func (f *RateLimiterGetTokenFunc) nextHook() func(context.Context, string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RateLimiterGetTokenFunc) appendCall(r0 RateLimiterGetTokenFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RateLimiterGetTokenFuncCall objects
// describing the invocations of this function.
func (f *RateLimiterGetTokenFunc) History() []RateLimiterGetTokenFuncCall {
	f.mutex.Lock()
	history := make([]RateLimiterGetTokenFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RateLimiterGetTokenFuncCall is an object that describes an invocation of
// method GetToken on an instance of MockRateLimiter.
type RateLimiterGetTokenFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RateLimiterGetTokenFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RateLimiterGetTokenFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// RateLimiterSetTokenBucketConfigFunc describes the behavior when the
// SetTokenBucketConfig method of the parent MockRateLimiter instance is
// invoked.
type RateLimiterSetTokenBucketConfigFunc struct {
	defaultHook func(context.Context, string, int32, time.Duration) error
	hooks       []func(context.Context, string, int32, time.Duration) error
	history     []RateLimiterSetTokenBucketConfigFuncCall
	mutex       sync.Mutex
}

// SetTokenBucketConfig delegates to the next hook function in the queue and
// stores the parameter and result values of this invocation.
func (m *MockRateLimiter) SetTokenBucketConfig(v0 context.Context, v1 string, v2 int32, v3 time.Duration) error {
	r0 := m.SetTokenBucketConfigFunc.nextHook()(v0, v1, v2, v3)
	m.SetTokenBucketConfigFunc.appendCall(RateLimiterSetTokenBucketConfigFuncCall{v0, v1, v2, v3, r0})
	return r0
}

// SetDefaultHook sets function that is called when the SetTokenBucketConfig
// method of the parent MockRateLimiter instance is invoked and the hook
// queue is empty.
func (f *RateLimiterSetTokenBucketConfigFunc) SetDefaultHook(hook func(context.Context, string, int32, time.Duration) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// SetTokenBucketConfig method of the parent MockRateLimiter instance
// invokes the hook at the front of the queue and discards it. After the
// queue is empty, the default hook function is invoked for any future
// action.
func (f *RateLimiterSetTokenBucketConfigFunc) PushHook(hook func(context.Context, string, int32, time.Duration) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *RateLimiterSetTokenBucketConfigFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, string, int32, time.Duration) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *RateLimiterSetTokenBucketConfigFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, string, int32, time.Duration) error {
		return r0
	})
}

func (f *RateLimiterSetTokenBucketConfigFunc) nextHook() func(context.Context, string, int32, time.Duration) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *RateLimiterSetTokenBucketConfigFunc) appendCall(r0 RateLimiterSetTokenBucketConfigFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of RateLimiterSetTokenBucketConfigFuncCall
// objects describing the invocations of this function.
func (f *RateLimiterSetTokenBucketConfigFunc) History() []RateLimiterSetTokenBucketConfigFuncCall {
	f.mutex.Lock()
	history := make([]RateLimiterSetTokenBucketConfigFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// RateLimiterSetTokenBucketConfigFuncCall is an object that describes an
// invocation of method SetTokenBucketConfig on an instance of
// MockRateLimiter.
type RateLimiterSetTokenBucketConfigFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 int32
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 time.Duration
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c RateLimiterSetTokenBucketConfigFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c RateLimiterSetTokenBucketConfigFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}
