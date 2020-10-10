package usecase

import "SkeletonAPI/modules/customer/model"

// ResultUseCase data structure
type ResultUseCase struct {
	Result interface{}
	Error  error
}

type CustomerUseCase interface {
	CheckSaldo(accountNumber int) ResultUseCase
	TransferSaldo(req model.TransferRequest) ResultUseCase
}
