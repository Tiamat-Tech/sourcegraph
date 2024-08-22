// Code generated by go-mockgen 1.3.7; DO NOT EDIT.
//
// This file was generated by running `sg generate` (or `go-mockgen`) at the root of
// this repository. To add additional mocks to this or another package, add a new entry
// to the mockgen.yaml file in the root of this repository.

package store

import (
	"context"
	"sync"
)

// MockJobTokenStore is a mock implementation of the JobTokenStore interface
// (from the package
// github.com/sourcegraph/sourcegraph/internal/executor/store) used for unit
// testing.
type MockJobTokenStore struct {
	// CreateFunc is an instance of a mock function object controlling the
	// behavior of the method Create.
	CreateFunc *JobTokenStoreCreateFunc
	// DeleteFunc is an instance of a mock function object controlling the
	// behavior of the method Delete.
	DeleteFunc *JobTokenStoreDeleteFunc
	// ExistsFunc is an instance of a mock function object controlling the
	// behavior of the method Exists.
	ExistsFunc *JobTokenStoreExistsFunc
	// GetFunc is an instance of a mock function object controlling the
	// behavior of the method Get.
	GetFunc *JobTokenStoreGetFunc
	// GetByTokenFunc is an instance of a mock function object controlling
	// the behavior of the method GetByToken.
	GetByTokenFunc *JobTokenStoreGetByTokenFunc
	// RegenerateFunc is an instance of a mock function object controlling
	// the behavior of the method Regenerate.
	RegenerateFunc *JobTokenStoreRegenerateFunc
}

// NewMockJobTokenStore creates a new mock of the JobTokenStore interface.
// All methods return zero values for all results, unless overwritten.
func NewMockJobTokenStore() *MockJobTokenStore {
	return &MockJobTokenStore{
		CreateFunc: &JobTokenStoreCreateFunc{
			defaultHook: func(context.Context, int, string, string) (r0 string, r1 error) {
				return
			},
		},
		DeleteFunc: &JobTokenStoreDeleteFunc{
			defaultHook: func(context.Context, int, string) (r0 error) {
				return
			},
		},
		ExistsFunc: &JobTokenStoreExistsFunc{
			defaultHook: func(context.Context, int, string) (r0 bool, r1 error) {
				return
			},
		},
		GetFunc: &JobTokenStoreGetFunc{
			defaultHook: func(context.Context, int, string) (r0 JobToken, r1 error) {
				return
			},
		},
		GetByTokenFunc: &JobTokenStoreGetByTokenFunc{
			defaultHook: func(context.Context, string) (r0 JobToken, r1 error) {
				return
			},
		},
		RegenerateFunc: &JobTokenStoreRegenerateFunc{
			defaultHook: func(context.Context, int, string) (r0 string, r1 error) {
				return
			},
		},
	}
}

// NewStrictMockJobTokenStore creates a new mock of the JobTokenStore
// interface. All methods panic on invocation, unless overwritten.
func NewStrictMockJobTokenStore() *MockJobTokenStore {
	return &MockJobTokenStore{
		CreateFunc: &JobTokenStoreCreateFunc{
			defaultHook: func(context.Context, int, string, string) (string, error) {
				panic("unexpected invocation of MockJobTokenStore.Create")
			},
		},
		DeleteFunc: &JobTokenStoreDeleteFunc{
			defaultHook: func(context.Context, int, string) error {
				panic("unexpected invocation of MockJobTokenStore.Delete")
			},
		},
		ExistsFunc: &JobTokenStoreExistsFunc{
			defaultHook: func(context.Context, int, string) (bool, error) {
				panic("unexpected invocation of MockJobTokenStore.Exists")
			},
		},
		GetFunc: &JobTokenStoreGetFunc{
			defaultHook: func(context.Context, int, string) (JobToken, error) {
				panic("unexpected invocation of MockJobTokenStore.Get")
			},
		},
		GetByTokenFunc: &JobTokenStoreGetByTokenFunc{
			defaultHook: func(context.Context, string) (JobToken, error) {
				panic("unexpected invocation of MockJobTokenStore.GetByToken")
			},
		},
		RegenerateFunc: &JobTokenStoreRegenerateFunc{
			defaultHook: func(context.Context, int, string) (string, error) {
				panic("unexpected invocation of MockJobTokenStore.Regenerate")
			},
		},
	}
}

// NewMockJobTokenStoreFrom creates a new mock of the MockJobTokenStore
// interface. All methods delegate to the given implementation, unless
// overwritten.
func NewMockJobTokenStoreFrom(i JobTokenStore) *MockJobTokenStore {
	return &MockJobTokenStore{
		CreateFunc: &JobTokenStoreCreateFunc{
			defaultHook: i.Create,
		},
		DeleteFunc: &JobTokenStoreDeleteFunc{
			defaultHook: i.Delete,
		},
		ExistsFunc: &JobTokenStoreExistsFunc{
			defaultHook: i.Exists,
		},
		GetFunc: &JobTokenStoreGetFunc{
			defaultHook: i.Get,
		},
		GetByTokenFunc: &JobTokenStoreGetByTokenFunc{
			defaultHook: i.GetByToken,
		},
		RegenerateFunc: &JobTokenStoreRegenerateFunc{
			defaultHook: i.Regenerate,
		},
	}
}

// JobTokenStoreCreateFunc describes the behavior when the Create method of
// the parent MockJobTokenStore instance is invoked.
type JobTokenStoreCreateFunc struct {
	defaultHook func(context.Context, int, string, string) (string, error)
	hooks       []func(context.Context, int, string, string) (string, error)
	history     []JobTokenStoreCreateFuncCall
	mutex       sync.Mutex
}

// Create delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockJobTokenStore) Create(v0 context.Context, v1 int, v2 string, v3 string) (string, error) {
	r0, r1 := m.CreateFunc.nextHook()(v0, v1, v2, v3)
	m.CreateFunc.appendCall(JobTokenStoreCreateFuncCall{v0, v1, v2, v3, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Create method of the
// parent MockJobTokenStore instance is invoked and the hook queue is empty.
func (f *JobTokenStoreCreateFunc) SetDefaultHook(hook func(context.Context, int, string, string) (string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Create method of the parent MockJobTokenStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *JobTokenStoreCreateFunc) PushHook(hook func(context.Context, int, string, string) (string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *JobTokenStoreCreateFunc) SetDefaultReturn(r0 string, r1 error) {
	f.SetDefaultHook(func(context.Context, int, string, string) (string, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *JobTokenStoreCreateFunc) PushReturn(r0 string, r1 error) {
	f.PushHook(func(context.Context, int, string, string) (string, error) {
		return r0, r1
	})
}

func (f *JobTokenStoreCreateFunc) nextHook() func(context.Context, int, string, string) (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *JobTokenStoreCreateFunc) appendCall(r0 JobTokenStoreCreateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of JobTokenStoreCreateFuncCall objects
// describing the invocations of this function.
func (f *JobTokenStoreCreateFunc) History() []JobTokenStoreCreateFuncCall {
	f.mutex.Lock()
	history := make([]JobTokenStoreCreateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// JobTokenStoreCreateFuncCall is an object that describes an invocation of
// method Create on an instance of MockJobTokenStore.
type JobTokenStoreCreateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Arg3 is the value of the 4th argument passed to this method
	// invocation.
	Arg3 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c JobTokenStoreCreateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2, c.Arg3}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c JobTokenStoreCreateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// JobTokenStoreDeleteFunc describes the behavior when the Delete method of
// the parent MockJobTokenStore instance is invoked.
type JobTokenStoreDeleteFunc struct {
	defaultHook func(context.Context, int, string) error
	hooks       []func(context.Context, int, string) error
	history     []JobTokenStoreDeleteFuncCall
	mutex       sync.Mutex
}

// Delete delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockJobTokenStore) Delete(v0 context.Context, v1 int, v2 string) error {
	r0 := m.DeleteFunc.nextHook()(v0, v1, v2)
	m.DeleteFunc.appendCall(JobTokenStoreDeleteFuncCall{v0, v1, v2, r0})
	return r0
}

// SetDefaultHook sets function that is called when the Delete method of the
// parent MockJobTokenStore instance is invoked and the hook queue is empty.
func (f *JobTokenStoreDeleteFunc) SetDefaultHook(hook func(context.Context, int, string) error) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Delete method of the parent MockJobTokenStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *JobTokenStoreDeleteFunc) PushHook(hook func(context.Context, int, string) error) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *JobTokenStoreDeleteFunc) SetDefaultReturn(r0 error) {
	f.SetDefaultHook(func(context.Context, int, string) error {
		return r0
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *JobTokenStoreDeleteFunc) PushReturn(r0 error) {
	f.PushHook(func(context.Context, int, string) error {
		return r0
	})
}

func (f *JobTokenStoreDeleteFunc) nextHook() func(context.Context, int, string) error {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *JobTokenStoreDeleteFunc) appendCall(r0 JobTokenStoreDeleteFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of JobTokenStoreDeleteFuncCall objects
// describing the invocations of this function.
func (f *JobTokenStoreDeleteFunc) History() []JobTokenStoreDeleteFuncCall {
	f.mutex.Lock()
	history := make([]JobTokenStoreDeleteFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// JobTokenStoreDeleteFuncCall is an object that describes an invocation of
// method Delete on an instance of MockJobTokenStore.
type JobTokenStoreDeleteFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c JobTokenStoreDeleteFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c JobTokenStoreDeleteFuncCall) Results() []interface{} {
	return []interface{}{c.Result0}
}

// JobTokenStoreExistsFunc describes the behavior when the Exists method of
// the parent MockJobTokenStore instance is invoked.
type JobTokenStoreExistsFunc struct {
	defaultHook func(context.Context, int, string) (bool, error)
	hooks       []func(context.Context, int, string) (bool, error)
	history     []JobTokenStoreExistsFuncCall
	mutex       sync.Mutex
}

// Exists delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockJobTokenStore) Exists(v0 context.Context, v1 int, v2 string) (bool, error) {
	r0, r1 := m.ExistsFunc.nextHook()(v0, v1, v2)
	m.ExistsFunc.appendCall(JobTokenStoreExistsFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Exists method of the
// parent MockJobTokenStore instance is invoked and the hook queue is empty.
func (f *JobTokenStoreExistsFunc) SetDefaultHook(hook func(context.Context, int, string) (bool, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Exists method of the parent MockJobTokenStore instance invokes the hook
// at the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *JobTokenStoreExistsFunc) PushHook(hook func(context.Context, int, string) (bool, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *JobTokenStoreExistsFunc) SetDefaultReturn(r0 bool, r1 error) {
	f.SetDefaultHook(func(context.Context, int, string) (bool, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *JobTokenStoreExistsFunc) PushReturn(r0 bool, r1 error) {
	f.PushHook(func(context.Context, int, string) (bool, error) {
		return r0, r1
	})
}

func (f *JobTokenStoreExistsFunc) nextHook() func(context.Context, int, string) (bool, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *JobTokenStoreExistsFunc) appendCall(r0 JobTokenStoreExistsFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of JobTokenStoreExistsFuncCall objects
// describing the invocations of this function.
func (f *JobTokenStoreExistsFunc) History() []JobTokenStoreExistsFuncCall {
	f.mutex.Lock()
	history := make([]JobTokenStoreExistsFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// JobTokenStoreExistsFuncCall is an object that describes an invocation of
// method Exists on an instance of MockJobTokenStore.
type JobTokenStoreExistsFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 bool
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c JobTokenStoreExistsFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c JobTokenStoreExistsFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// JobTokenStoreGetFunc describes the behavior when the Get method of the
// parent MockJobTokenStore instance is invoked.
type JobTokenStoreGetFunc struct {
	defaultHook func(context.Context, int, string) (JobToken, error)
	hooks       []func(context.Context, int, string) (JobToken, error)
	history     []JobTokenStoreGetFuncCall
	mutex       sync.Mutex
}

// Get delegates to the next hook function in the queue and stores the
// parameter and result values of this invocation.
func (m *MockJobTokenStore) Get(v0 context.Context, v1 int, v2 string) (JobToken, error) {
	r0, r1 := m.GetFunc.nextHook()(v0, v1, v2)
	m.GetFunc.appendCall(JobTokenStoreGetFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Get method of the
// parent MockJobTokenStore instance is invoked and the hook queue is empty.
func (f *JobTokenStoreGetFunc) SetDefaultHook(hook func(context.Context, int, string) (JobToken, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Get method of the parent MockJobTokenStore instance invokes the hook at
// the front of the queue and discards it. After the queue is empty, the
// default hook function is invoked for any future action.
func (f *JobTokenStoreGetFunc) PushHook(hook func(context.Context, int, string) (JobToken, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *JobTokenStoreGetFunc) SetDefaultReturn(r0 JobToken, r1 error) {
	f.SetDefaultHook(func(context.Context, int, string) (JobToken, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *JobTokenStoreGetFunc) PushReturn(r0 JobToken, r1 error) {
	f.PushHook(func(context.Context, int, string) (JobToken, error) {
		return r0, r1
	})
}

func (f *JobTokenStoreGetFunc) nextHook() func(context.Context, int, string) (JobToken, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *JobTokenStoreGetFunc) appendCall(r0 JobTokenStoreGetFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of JobTokenStoreGetFuncCall objects describing
// the invocations of this function.
func (f *JobTokenStoreGetFunc) History() []JobTokenStoreGetFuncCall {
	f.mutex.Lock()
	history := make([]JobTokenStoreGetFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// JobTokenStoreGetFuncCall is an object that describes an invocation of
// method Get on an instance of MockJobTokenStore.
type JobTokenStoreGetFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 JobToken
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c JobTokenStoreGetFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c JobTokenStoreGetFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// JobTokenStoreGetByTokenFunc describes the behavior when the GetByToken
// method of the parent MockJobTokenStore instance is invoked.
type JobTokenStoreGetByTokenFunc struct {
	defaultHook func(context.Context, string) (JobToken, error)
	hooks       []func(context.Context, string) (JobToken, error)
	history     []JobTokenStoreGetByTokenFuncCall
	mutex       sync.Mutex
}

// GetByToken delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockJobTokenStore) GetByToken(v0 context.Context, v1 string) (JobToken, error) {
	r0, r1 := m.GetByTokenFunc.nextHook()(v0, v1)
	m.GetByTokenFunc.appendCall(JobTokenStoreGetByTokenFuncCall{v0, v1, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the GetByToken method of
// the parent MockJobTokenStore instance is invoked and the hook queue is
// empty.
func (f *JobTokenStoreGetByTokenFunc) SetDefaultHook(hook func(context.Context, string) (JobToken, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// GetByToken method of the parent MockJobTokenStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *JobTokenStoreGetByTokenFunc) PushHook(hook func(context.Context, string) (JobToken, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *JobTokenStoreGetByTokenFunc) SetDefaultReturn(r0 JobToken, r1 error) {
	f.SetDefaultHook(func(context.Context, string) (JobToken, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *JobTokenStoreGetByTokenFunc) PushReturn(r0 JobToken, r1 error) {
	f.PushHook(func(context.Context, string) (JobToken, error) {
		return r0, r1
	})
}

func (f *JobTokenStoreGetByTokenFunc) nextHook() func(context.Context, string) (JobToken, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *JobTokenStoreGetByTokenFunc) appendCall(r0 JobTokenStoreGetByTokenFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of JobTokenStoreGetByTokenFuncCall objects
// describing the invocations of this function.
func (f *JobTokenStoreGetByTokenFunc) History() []JobTokenStoreGetByTokenFuncCall {
	f.mutex.Lock()
	history := make([]JobTokenStoreGetByTokenFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// JobTokenStoreGetByTokenFuncCall is an object that describes an invocation
// of method GetByToken on an instance of MockJobTokenStore.
type JobTokenStoreGetByTokenFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 JobToken
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c JobTokenStoreGetByTokenFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c JobTokenStoreGetByTokenFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}

// JobTokenStoreRegenerateFunc describes the behavior when the Regenerate
// method of the parent MockJobTokenStore instance is invoked.
type JobTokenStoreRegenerateFunc struct {
	defaultHook func(context.Context, int, string) (string, error)
	hooks       []func(context.Context, int, string) (string, error)
	history     []JobTokenStoreRegenerateFuncCall
	mutex       sync.Mutex
}

// Regenerate delegates to the next hook function in the queue and stores
// the parameter and result values of this invocation.
func (m *MockJobTokenStore) Regenerate(v0 context.Context, v1 int, v2 string) (string, error) {
	r0, r1 := m.RegenerateFunc.nextHook()(v0, v1, v2)
	m.RegenerateFunc.appendCall(JobTokenStoreRegenerateFuncCall{v0, v1, v2, r0, r1})
	return r0, r1
}

// SetDefaultHook sets function that is called when the Regenerate method of
// the parent MockJobTokenStore instance is invoked and the hook queue is
// empty.
func (f *JobTokenStoreRegenerateFunc) SetDefaultHook(hook func(context.Context, int, string) (string, error)) {
	f.defaultHook = hook
}

// PushHook adds a function to the end of hook queue. Each invocation of the
// Regenerate method of the parent MockJobTokenStore instance invokes the
// hook at the front of the queue and discards it. After the queue is empty,
// the default hook function is invoked for any future action.
func (f *JobTokenStoreRegenerateFunc) PushHook(hook func(context.Context, int, string) (string, error)) {
	f.mutex.Lock()
	f.hooks = append(f.hooks, hook)
	f.mutex.Unlock()
}

// SetDefaultReturn calls SetDefaultHook with a function that returns the
// given values.
func (f *JobTokenStoreRegenerateFunc) SetDefaultReturn(r0 string, r1 error) {
	f.SetDefaultHook(func(context.Context, int, string) (string, error) {
		return r0, r1
	})
}

// PushReturn calls PushHook with a function that returns the given values.
func (f *JobTokenStoreRegenerateFunc) PushReturn(r0 string, r1 error) {
	f.PushHook(func(context.Context, int, string) (string, error) {
		return r0, r1
	})
}

func (f *JobTokenStoreRegenerateFunc) nextHook() func(context.Context, int, string) (string, error) {
	f.mutex.Lock()
	defer f.mutex.Unlock()

	if len(f.hooks) == 0 {
		return f.defaultHook
	}

	hook := f.hooks[0]
	f.hooks = f.hooks[1:]
	return hook
}

func (f *JobTokenStoreRegenerateFunc) appendCall(r0 JobTokenStoreRegenerateFuncCall) {
	f.mutex.Lock()
	f.history = append(f.history, r0)
	f.mutex.Unlock()
}

// History returns a sequence of JobTokenStoreRegenerateFuncCall objects
// describing the invocations of this function.
func (f *JobTokenStoreRegenerateFunc) History() []JobTokenStoreRegenerateFuncCall {
	f.mutex.Lock()
	history := make([]JobTokenStoreRegenerateFuncCall, len(f.history))
	copy(history, f.history)
	f.mutex.Unlock()

	return history
}

// JobTokenStoreRegenerateFuncCall is an object that describes an invocation
// of method Regenerate on an instance of MockJobTokenStore.
type JobTokenStoreRegenerateFuncCall struct {
	// Arg0 is the value of the 1st argument passed to this method
	// invocation.
	Arg0 context.Context
	// Arg1 is the value of the 2nd argument passed to this method
	// invocation.
	Arg1 int
	// Arg2 is the value of the 3rd argument passed to this method
	// invocation.
	Arg2 string
	// Result0 is the value of the 1st result returned from this method
	// invocation.
	Result0 string
	// Result1 is the value of the 2nd result returned from this method
	// invocation.
	Result1 error
}

// Args returns an interface slice containing the arguments of this
// invocation.
func (c JobTokenStoreRegenerateFuncCall) Args() []interface{} {
	return []interface{}{c.Arg0, c.Arg1, c.Arg2}
}

// Results returns an interface slice containing the results of this
// invocation.
func (c JobTokenStoreRegenerateFuncCall) Results() []interface{} {
	return []interface{}{c.Result0, c.Result1}
}
