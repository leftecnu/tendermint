// Code generated by mockery. DO NOT EDIT.

package mocks

import (
	context "context"

	mock "github.com/stretchr/testify/mock"

	state "github.com/leftecnu/tendermint/state"

	types "github.com/leftecnu/tendermint/types"
)

// StateProvider is an autogenerated mock type for the StateProvider type
type StateProvider struct {
	mock.Mock
}

// AppHash provides a mock function with given fields: ctx, height
func (_m *StateProvider) AppHash(ctx context.Context, height uint64) ([]byte, error) {
	ret := _m.Called(ctx, height)

	var r0 []byte
	if rf, ok := ret.Get(0).(func(context.Context, uint64) []byte); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Commit provides a mock function with given fields: ctx, height
func (_m *StateProvider) Commit(ctx context.Context, height uint64) (*types.Commit, error) {
	ret := _m.Called(ctx, height)

	var r0 *types.Commit
	if rf, ok := ret.Get(0).(func(context.Context, uint64) *types.Commit); ok {
		r0 = rf(ctx, height)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*types.Commit)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// State provides a mock function with given fields: ctx, height
func (_m *StateProvider) State(ctx context.Context, height uint64) (state.State, error) {
	ret := _m.Called(ctx, height)

	var r0 state.State
	if rf, ok := ret.Get(0).(func(context.Context, uint64) state.State); ok {
		r0 = rf(ctx, height)
	} else {
		r0 = ret.Get(0).(state.State)
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(context.Context, uint64) error); ok {
		r1 = rf(ctx, height)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewStateProviderT interface {
	mock.TestingT
	Cleanup(func())
}

// NewStateProvider creates a new instance of StateProvider. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewStateProvider(t NewStateProviderT) *StateProvider {
	mock := &StateProvider{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
