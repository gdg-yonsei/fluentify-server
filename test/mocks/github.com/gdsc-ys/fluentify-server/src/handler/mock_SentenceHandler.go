// Code generated by mockery v2.40.3. DO NOT EDIT.

package handler_test

import (
	echo "github.com/labstack/echo/v4"

	mock "github.com/stretchr/testify/mock"
)

// MockSentenceHandler is an autogenerated mock type for the SentenceHandler type
type MockSentenceHandler struct {
	mock.Mock
}

type MockSentenceHandler_Expecter struct {
	mock *mock.Mock
}

func (_m *MockSentenceHandler) EXPECT() *MockSentenceHandler_Expecter {
	return &MockSentenceHandler_Expecter{mock: &_m.Mock}
}

// GetSentence provides a mock function with given fields: c
func (_m *MockSentenceHandler) GetSentence(c echo.Context) error {
	ret := _m.Called(c)

	if len(ret) == 0 {
		panic("no return value specified for GetSentence")
	}

	var r0 error
	if rf, ok := ret.Get(0).(func(echo.Context) error); ok {
		r0 = rf(c)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// MockSentenceHandler_GetSentence_Call is a *mock.Call that shadows Run/Return methods with type explicit version for method 'GetSentence'
type MockSentenceHandler_GetSentence_Call struct {
	*mock.Call
}

// GetSentence is a helper method to define mock.On call
//   - c echo.Context
func (_e *MockSentenceHandler_Expecter) GetSentence(c interface{}) *MockSentenceHandler_GetSentence_Call {
	return &MockSentenceHandler_GetSentence_Call{Call: _e.mock.On("GetSentence", c)}
}

func (_c *MockSentenceHandler_GetSentence_Call) Run(run func(c echo.Context)) *MockSentenceHandler_GetSentence_Call {
	_c.Call.Run(func(args mock.Arguments) {
		run(args[0].(echo.Context))
	})
	return _c
}

func (_c *MockSentenceHandler_GetSentence_Call) Return(_a0 error) *MockSentenceHandler_GetSentence_Call {
	_c.Call.Return(_a0)
	return _c
}

func (_c *MockSentenceHandler_GetSentence_Call) RunAndReturn(run func(echo.Context) error) *MockSentenceHandler_GetSentence_Call {
	_c.Call.Return(run)
	return _c
}

// NewMockSentenceHandler creates a new instance of MockSentenceHandler. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
// The first argument is typically a *testing.T value.
func NewMockSentenceHandler(t interface {
	mock.TestingT
	Cleanup(func())
}) *MockSentenceHandler {
	mock := &MockSentenceHandler{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}