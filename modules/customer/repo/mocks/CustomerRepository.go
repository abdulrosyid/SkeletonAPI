// Code generated by mockery v2.2.1. DO NOT EDIT.

package mocks

import (
	repo "SkeletonAPI/modules/customer/repo"

	mock "github.com/stretchr/testify/mock"
)

// CustomerRepository is an autogenerated mock type for the CustomerRepository type
type CustomerRepository struct {
	mock.Mock
}

// GetCustomerSaldo provides a mock function with given fields: accountNumber
func (_m *CustomerRepository) GetCustomerSaldo(accountNumber int) repo.ResultRepository {
	ret := _m.Called(accountNumber)

	var r0 repo.ResultRepository
	if rf, ok := ret.Get(0).(func(int) repo.ResultRepository); ok {
		r0 = rf(accountNumber)
	} else {
		r0 = ret.Get(0).(repo.ResultRepository)
	}

	return r0
}

// TransferSaldo provides a mock function with given fields: fromAccountNumber, toAccountNumber, balance
func (_m *CustomerRepository) TransferSaldo(fromAccountNumber int, toAccountNumber int, balance int) repo.ResultRepository {
	ret := _m.Called(fromAccountNumber, toAccountNumber, balance)

	var r0 repo.ResultRepository
	if rf, ok := ret.Get(0).(func(int, int, int) repo.ResultRepository); ok {
		r0 = rf(fromAccountNumber, toAccountNumber, balance)
	} else {
		r0 = ret.Get(0).(repo.ResultRepository)
	}

	return r0
}
