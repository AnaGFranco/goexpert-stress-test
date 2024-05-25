// Code generated by mockery v2.43.0. DO NOT EDIT.

package mock

import (
	domain "github.com/anagfranco/goexpert-stress-test/internal/domain"
	mock "github.com/stretchr/testify/mock"
)

// LoadTester is an autogenerated mock type for the LoadTester type
type LoadTester struct {
	mock.Mock
}

type LoadTester_Expecter struct {
	mock *mock.Mock
}

func (_m *LoadTester) EXPECT() *LoadTester_Expecter {
	return &LoadTester_Expecter{mock: &_m.Mock}
}

// RunLoadTest provides a mock function with given fields: config
func (_m *LoadTester) RunLoadTest(config domain.TestConfig) domain.TestResult {
	ret := _m.Called(config)

	if len(ret) == 0 {
		panic("no return value specified for RunLoadTest")
	}

	var r0 domain.TestResult
	if rf, ok := ret.Get(0).(func(domain.TestConfig) domain.TestResult); ok {
		r0 = rf(config)
	} else {
		r0 = ret.Get(0).(domain.TestResult)
	}

	return r0
}

// LoadTester_RunLoadTest_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'RunLoadTest'
type LoadTester_RunLoadTest_Call struct {
	*mock.Call
}

// RunLoadTest is a helper method to define mock.On call
//   - config domain.TestConfig
func (_e *LoadTester_Expecter) RunLoadTest(config interface{}) *LoadTester_RunLoadTest_Call {
	return &LoadTester_RunLoadTest_Call{Call: _e.mock.On("RunLoadTest", config)}
}

func (_c *LoadTester_RunLoadTest_Call) Run(run func(config domain.TestConfig)) *LoadTester_RunLoadTest_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(domain.TestConfig))
	})
	return _c
}

func (_c *LoadTester_RunLoadTest_Call) Return(_a0 domain.TestResult) *LoadTester_RunLoadTest_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *LoadTester_RunLoadTest_Call) RunAndReturn(run func(domain.TestConfig) domain.TestResult) *LoadTester_RunLoadTest_Call {
	_c.Call.Return(run)
	return _c
}

// NewLoadTester creates a new instance of LoadTester. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewLoadTester(t interface {
	mock.TestingT
	Cleanup(func())
}) *LoadTester {
	mock := &LoadTester{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
