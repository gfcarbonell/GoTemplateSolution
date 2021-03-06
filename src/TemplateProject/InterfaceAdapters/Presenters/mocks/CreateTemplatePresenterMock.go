// Code generated by mockery v2.12.3. DO NOT EDIT.

package mockPresenters

import mock "github.com/stretchr/testify/mock"

// ICreateTemplatePresenter is an autogenerated mock type for the ICreateTemplatePresenter type
type ICreateTemplatePresenter struct {
	mock.Mock
}

// GetContent provides a mock function with given fields:
func (_m *ICreateTemplatePresenter) GetContent() int64 {
	ret := _m.Called()

	var r0 int64
	if rf, ok := ret.Get(0).(func() int64); ok {
		r0 = rf()
	} else {
		r0 = ret.Get(0).(int64)
	}

	return r0
}

// GetError provides a mock function with given fields:
func (_m *ICreateTemplatePresenter) GetError() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Handle provides a mock function with given fields: response, err
func (_m *ICreateTemplatePresenter) Handle(response int64, err error) {
	_m.Called(response, err)
}

type NewICreateTemplatePresenterT interface {
	mock.TestingT
	Cleanup(func())
}

// NewICreateTemplatePresenter creates a new instance of ICreateTemplatePresenter. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICreateTemplatePresenter(t NewICreateTemplatePresenterT) *ICreateTemplatePresenter {
	mock := &ICreateTemplatePresenter{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}
