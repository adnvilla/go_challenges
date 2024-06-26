// Code generated by mockery v2.42.2. DO NOT EDIT.

package mock

import mock "github.com/stretchr/testify/mock"

// MockCreditAssigner is an autogenerated mock type for the CreditAssigner type
type MockCreditAssigner struct {
	mock.Mock
}

type MockCreditAssigner_Expecter struct {
	mock *mock.Mock
}

func (_m *MockCreditAssigner) EXPECT() *MockCreditAssigner_Expecter {
	return &MockCreditAssigner_Expecter{mock: &_m.Mock}
}

// Assign provides a mock function with given fields: investment
func (_m *MockCreditAssigner) Assign(investment int32) (int32, int32, int32, error) {
	ret := _m.Called(investment)

	if len(ret) == 0 {
		panic("no return value specified for Assign")
	}

	var r0 int32
	var r1 int32
	var r2 int32
	var r3 error
	if rf, ok := ret.Get(0).(func(int32) (int32, int32, int32, error)); ok {
		return rf(investment)
	}
	if rf, ok := ret.Get(0).(func(int32) int32); ok {
		r0 = rf(investment)
	} else {
		r0 = ret.Get(0).(int32)
	}

	if rf, ok := ret.Get(1).(func(int32) int32); ok {
		r1 = rf(investment)
	} else {
		r1 = ret.Get(1).(int32)
	}

	if rf, ok := ret.Get(2).(func(int32) int32); ok {
		r2 = rf(investment)
	} else {
		r2 = ret.Get(2).(int32)
	}

	if rf, ok := ret.Get(3).(func(int32) error); ok {
		r3 = rf(investment)
	} else {
		r3 = ret.Error(3)
	}

	return r0, r1, r2, r3
}

// MockCreditAssigner_Assign_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'Assign'
type MockCreditAssigner_Assign_Call struct {
	*mock.Call
}

// Assign is a helper method to define mock.On call
//   - investment int32
func (_e *MockCreditAssigner_Expecter) Assign(investment interface{}) *MockCreditAssigner_Assign_Call {
	return &MockCreditAssigner_Assign_Call{Call: _e.mock.On("Assign", investment)}
}

func (_c *MockCreditAssigner_Assign_Call) Run(run func(investment int32)) *MockCreditAssigner_Assign_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(int32))
	})
	return _c
}

func (_c *MockCreditAssigner_Assign_Call) Return(_a0 int32, _a1 int32, _a2 int32, _a3 error) *MockCreditAssigner_Assign_Call {
	_c.Call.Return(_a0, _a1, _a2, _a3)
	return _c
}

func (_c *MockCreditAssigner_Assign_Call) RunAndReturn(run func(int32) (int32, int32, int32, error)) *MockCreditAssigner_Assign_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockCreditAssigner creates a new instance of MockCreditAssigner. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockCreditAssigner(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockCreditAssigner {
	mock := &MockCreditAssigner{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
