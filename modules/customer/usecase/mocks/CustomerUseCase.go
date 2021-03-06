// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	model "SkeletonAPI/modules/customer/model"

	mock "github.com/stretchr/testify/mock"

	usecase "SkeletonAPI/modules/customer/usecase"
)

// CustomerUseCase is an autogenerated mock type for the CustomerUseCase type
type CustomerUseCase struct {
	mock.Mock
}

// CheckSaldo provides a mock function with given fields: accountNumber
func (_m *CustomerUseCase) CheckSaldo(accountNumber int) usecase.ResultUseCase {
	ret := _m.Called(accountNumber)

	var r0 usecase.ResultUseCase
	if rf, ok := ret.Get(0).(func(int) usecase.ResultUseCase); ok {
		r0 = rf(accountNumber)
	} else {
		r0 = ret.Get(0).(usecase.ResultUseCase)
	}

	return r0
}

// TransferSaldo provides a mock function with given fields: req
func (_m *CustomerUseCase) TransferSaldo(req model.TransferRequest) usecase.ResultUseCase {
	ret := _m.Called(req)

	var r0 usecase.ResultUseCase
	if rf, ok := ret.Get(0).(func(model.TransferRequest) usecase.ResultUseCase); ok {
		r0 = rf(req)
	} else {
		r0 = ret.Get(0).(usecase.ResultUseCase)
	}

	return r0
}
