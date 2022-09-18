// Code generated by mockery v2.14.0. DO NOT EDIT.

package mocks

import (
	domain "discountapp/domain"
	inputs "discountapp/usecases/inputs"

	mock "github.com/stretchr/testify/mock"
)

// IGetProduct is an autogenerated mock type for the IGetProduct type
type IGetProduct struct {
	mock.Mock
}

// Perform provides a mock function with given fields: productInput
func (_m *IGetProduct) Perform(productInput *inputs.ProductBySlugInput) (*domain.Product, error) {
	ret := _m.Called(productInput)

	var r0 *domain.Product
	if rf, ok := ret.Get(0).(func(*inputs.ProductBySlugInput) *domain.Product); ok {
		r0 = rf(productInput)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*domain.Product)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func(*inputs.ProductBySlugInput) error); ok {
		r1 = rf(productInput)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

type mockConstructorTestingTNewIGetProduct interface {
	mock.TestingT
	Cleanup(func())
}

// NewIGetProduct creates a new instance of IGetProduct. It also registers a testing interface on the mock and a cleanup function to assert the mocks expectations.
func NewIGetProduct(t mockConstructorTestingTNewIGetProduct) *IGetProduct {
	mock := &IGetProduct{}
	mock.Mock.Test(t)

	t.Cleanup(func() { mock.AssertExpectations(t) })

	return mock
}