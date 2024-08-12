// Code generated by mockery v2.43.2. DO NOT EDIT.

package blockchain

import (
	context "context"
	big "math/big"

	ethereum "github.com/ethereum/go-ethereum"

	mock "github.com/stretchr/testify/mock"

	types "github.com/ethereum/go-ethereum/core/types"
)

// MockChainClient is an autogenerated mock type for the ChainClient type
type MockChainClient struct {
	mock.Mock
}

type MockChainClient_Expecter struct {
	mock *mock.Mock
}

func (_m *MockChainClient) EXPECT() *MockChainClient_Expecter {
	return &MockChainClient_Expecter{mock: &_m.Mock}
}

// BlockNumber provides a mock function with given fields: ctx
func (_m *MockChainClient) BlockNumber(ctx context.Context) (uint64, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for BlockNumber")
	}

	var r0 uint64
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (uint64, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) uint64); ok {
		r0 = rf(ctx)
	} else {
		r0 = ret.Get(0).(uint64)
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainClient_BlockNumber_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'BlockNumber'
type MockChainClient_BlockNumber_Call struct {
	*mock.Call
}

// BlockNumber is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockChainClient_Expecter) BlockNumber(ctx interface{}) *MockChainClient_BlockNumber_Call {
	return &MockChainClient_BlockNumber_Call{Call: _e.mock.On("BlockNumber", ctx)}
}

func (_c *MockChainClient_BlockNumber_Call) Run(run func(ctx context.Context)) *MockChainClient_BlockNumber_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockChainClient_BlockNumber_Call) Return(_a0 uint64, _a1 error) *MockChainClient_BlockNumber_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainClient_BlockNumber_Call) RunAndReturn(run func(context.Context) (uint64, error)) *MockChainClient_BlockNumber_Call {
	_c.Call.Return(run)
	return _c
}

// ChainID provides a mock function with given fields: ctx
func (_m *MockChainClient) ChainID(ctx context.Context) (*big.Int, error) {
	ret := _m.Called(ctx)

	if len(ret) == 0 {
		panic("no return value specified for ChainID")
	}

	var r0 *big.Int
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context) (*big.Int, error)); ok {
		return rf(ctx)
	}
	if rf, ok := ret.Get(0).(func(context.Context) *big.Int); ok {
		r0 = rf(ctx)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*big.Int)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context) error); ok {
		r1 = rf(ctx)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainClient_ChainID_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'ChainID'
type MockChainClient_ChainID_Call struct {
	*mock.Call
}

// ChainID is a helper method to define mock.On call
//   - ctx context.Context
func (_e *MockChainClient_Expecter) ChainID(ctx interface{}) *MockChainClient_ChainID_Call {
	return &MockChainClient_ChainID_Call{Call: _e.mock.On("ChainID", ctx)}
}

func (_c *MockChainClient_ChainID_Call) Run(run func(ctx context.Context)) *MockChainClient_ChainID_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context))
	})
	return _c
}

func (_c *MockChainClient_ChainID_Call) Return(_a0 *big.Int, _a1 error) *MockChainClient_ChainID_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainClient_ChainID_Call) RunAndReturn(run func(context.Context) (*big.Int, error)) *MockChainClient_ChainID_Call {
	_c.Call.Return(run)
	return _c
}

// FilterLogs provides a mock function with given fields: ctx, q
func (_m *MockChainClient) FilterLogs(ctx context.Context, q ethereum.FilterQuery) ([]types.Log, error) {
	ret := _m.Called(ctx, q)

	if len(ret) == 0 {
		panic("no return value specified for FilterLogs")
	}

	var r0 []types.Log
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) ([]types.Log, error)); ok {
		return rf(ctx, q)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery) []types.Log); ok {
		r0 = rf(ctx, q)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]types.Log)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery) error); ok {
		r1 = rf(ctx, q)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainClient_FilterLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'FilterLogs'
type MockChainClient_FilterLogs_Call struct {
	*mock.Call
}

// FilterLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - q ethereum.FilterQuery
func (_e *MockChainClient_Expecter) FilterLogs(ctx interface{}, q interface{}) *MockChainClient_FilterLogs_Call {
	return &MockChainClient_FilterLogs_Call{Call: _e.mock.On("FilterLogs", ctx, q)}
}

func (_c *MockChainClient_FilterLogs_Call) Run(run func(ctx context.Context, q ethereum.FilterQuery)) *MockChainClient_FilterLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.FilterQuery))
	})
	return _c
}

func (_c *MockChainClient_FilterLogs_Call) Return(_a0 []types.Log, _a1 error) *MockChainClient_FilterLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainClient_FilterLogs_Call) RunAndReturn(run func(context.Context, ethereum.FilterQuery) ([]types.Log, error)) *MockChainClient_FilterLogs_Call {
	_c.Call.Return(run)
	return _c
}

// SubscribeFilterLogs provides a mock function with given fields: ctx, q, ch
func (_m *MockChainClient) SubscribeFilterLogs(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	ret := _m.Called(ctx, q, ch)

	if len(ret) == 0 {
		panic("no return value specified for SubscribeFilterLogs")
	}

	var r0 ethereum.Subscription
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)); ok {
		return rf(ctx, q, ch)
	}
	if rf, ok := ret.Get(0).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) ethereum.Subscription); ok {
		r0 = rf(ctx, q, ch)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(ethereum.Subscription)
		}
	}

	if rf, ok := ret.Get(1).(func(context.Context, ethereum.FilterQuery, chan<- types.Log) error); ok {
		r1 = rf(ctx, q, ch)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockChainClient_SubscribeFilterLogs_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'SubscribeFilterLogs'
type MockChainClient_SubscribeFilterLogs_Call struct {
	*mock.Call
}

// SubscribeFilterLogs is a helper method to define mock.On call
//   - ctx context.Context
//   - q ethereum.FilterQuery
//   - ch chan<- types.Log
func (_e *MockChainClient_Expecter) SubscribeFilterLogs(ctx interface{}, q interface{}, ch interface{}) *MockChainClient_SubscribeFilterLogs_Call {
	return &MockChainClient_SubscribeFilterLogs_Call{Call: _e.mock.On("SubscribeFilterLogs", ctx, q, ch)}
}

func (_c *MockChainClient_SubscribeFilterLogs_Call) Run(run func(ctx context.Context, q ethereum.FilterQuery, ch chan<- types.Log)) *MockChainClient_SubscribeFilterLogs_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(ethereum.FilterQuery), args[2].(chan<- types.Log))
	})
	return _c
}

func (_c *MockChainClient_SubscribeFilterLogs_Call) Return(_a0 ethereum.Subscription, _a1 error) *MockChainClient_SubscribeFilterLogs_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockChainClient_SubscribeFilterLogs_Call) RunAndReturn(run func(context.Context, ethereum.FilterQuery, chan<- types.Log) (ethereum.Subscription, error)) *MockChainClient_SubscribeFilterLogs_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockChainClient creates a new instance of MockChainClient. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockChainClient(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockChainClient {
	mock := &MockChainClient{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
