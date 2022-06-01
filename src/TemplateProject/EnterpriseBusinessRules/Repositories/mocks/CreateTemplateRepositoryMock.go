// Code generated by mockery v2.12.3. DO NOT EDIT.

package mockRepositories

import (
	Entities "TemplateSolution/src/TemplateProject/EnterpriseBusinessRules/Entities"

	mock "github.com/stretchr/testify/mock"
)

// ICreateTemplateRepository is an autogenerated mock type for the ICreateTemplateRepository type
type ICreateTemplateRepository struct {
	mock.Mock
}

// Handle provides a mock function with given fields: entity
func (_m *ICreateTemplateRepository) Handle(entity Entities.TemplateEntity) (*int64, error) {
	ret := _m.Called(entity)

	var r0 *int64
	if rf, ok := ret.Get(0).(func(Entities.TemplateEntity) *int64); ok {
		r0 = rf(entity)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*int64)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(Entities.TemplateEntity) error); ok {
		r1 = rf(entity)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type NewICreateTemplateRepositoryT interface {
	mock.TestingT
	Cleanup(func())
}

// NewICreateTemplateRepository creates a new instance of ICreateTemplateRepository. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewICreateTemplateRepository(t NewICreateTemplateRepositoryT) *ICreateTemplateRepository {
	mock := &ICreateTemplateRepository{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}