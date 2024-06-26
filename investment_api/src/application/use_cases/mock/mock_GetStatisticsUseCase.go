// Code generated by mockery v2.42.2. DO NOT EDIT.

package mock

import (
	context "context"

	use_cases "github.com/adnvilla/go_challenges/investment_api/src/application/use_cases"
	mock "github.com/stretchr/testify/mock"
)

// MockGetStatisticsUseCase is an autogenerated mock type for the GetStatisticsUseCase type
type MockGetStatisticsUseCase struct {
	mock.Mock
}

type MockGetStatisticsUseCase_Expecter struct {
	mock *mock.Mock
}

func (_m *MockGetStatisticsUseCase) EXPECT() *MockGetStatisticsUseCase_Expecter {
	return &MockGetStatisticsUseCase_Expecter{mock: &_m.Mock}
}

// Handle provides a mock function with given fields: ctx, input
func (_m *MockGetStatisticsUseCase) Handle(ctx context.Context, input use_cases.GetStatisticsInput) (use_cases.GetStatisticsOutput, error) {
	ret := _m.Called(ctx, input)

	if len(ret) == 0 {
		panic("no return value specified for Handle")
	}

	var r0 use_cases.GetStatisticsOutput
	var r1 error
	if rf, ok := ret.Get(0).(func(context.Context, use_cases.GetStatisticsInput) (use_cases.GetStatisticsOutput, error)); ok {
		return rf(ctx, input)
	}
	if rf, ok := ret.Get(0).(func(context.Context, use_cases.GetStatisticsInput) use_cases.GetStatisticsOutput); ok {
		r0 = rf(ctx, input)
	} else {
		r0 = ret.Get(0).(use_cases.GetStatisticsOutput)
	}

	if rf, ok := ret.Get(1).(func(context.Context, use_cases.GetStatisticsInput) error); ok {
		r1 = rf(ctx, input)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// MockGetStatisticsUseCase_Handle_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Handle'
type MockGetStatisticsUseCase_Handle_Call struct {
	*mock.Call
}

// Handle is a helper method to define mock.On call
//   - ctx context.Context
//   - input use_cases.GetStatisticsInput
func (_e *MockGetStatisticsUseCase_Expecter) Handle(ctx interface{}, input interface{}) *MockGetStatisticsUseCase_Handle_Call {
	return &MockGetStatisticsUseCase_Handle_Call{Call: _e.mock.On("Handle", ctx, input)}
}

func (_c *MockGetStatisticsUseCase_Handle_Call) Run(run func(ctx context.Context, input use_cases.GetStatisticsInput)) *MockGetStatisticsUseCase_Handle_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(context.Context), args[1].(use_cases.GetStatisticsInput))
	})
	return _c
}

func (_c *MockGetStatisticsUseCase_Handle_Call) Return(_a0 use_cases.GetStatisticsOutput, _a1 error) *MockGetStatisticsUseCase_Handle_Call {
	_c.Call.Return(_a0, _a1)
	return _c
}

func (_c *MockGetStatisticsUseCase_Handle_Call) RunAndReturn(run func(context.Context, use_cases.GetStatisticsInput) (use_cases.GetStatisticsOutput, error)) *MockGetStatisticsUseCase_Handle_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockGetStatisticsUseCase creates a new instance of MockGetStatisticsUseCase. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockGetStatisticsUseCase(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockGetStatisticsUseCase {
	mock := &MockGetStatisticsUseCase{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
