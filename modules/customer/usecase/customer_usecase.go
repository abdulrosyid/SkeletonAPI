package usecase

import (
	"SkeletonAPI/helper"
	"SkeletonAPI/modules/customer/model"
	"SkeletonAPI/modules/customer/repo"
	"database/sql"
	"fmt"
)

type CustomerUseCaseImpl struct {
	CustomerRepo repo.CustomerRepository
}

func NewCustomerUseCase(messageRepo repo.CustomerRepository) CustomerUseCase {
	return &CustomerUseCaseImpl{CustomerRepo: messageRepo}
}

func (c *CustomerUseCaseImpl) CheckSaldo(accountNumber int) ResultUseCase {
	result := ResultUseCase{}
	resultRepo := c.CustomerRepo.GetCustomerSaldo(accountNumber)
	if resultRepo.Error != nil {
		if resultRepo.Error == sql.ErrNoRows {
			result.Error = fmt.Errorf(helper.ErrorDataNotFound, "account")
			return result
		}
		result.Error = resultRepo.Error
		return result
	}

	result.Result = resultRepo.Result

	return result
}

func (c *CustomerUseCaseImpl) TransferSaldo(req model.TransferRequest) ResultUseCase {
	result := ResultUseCase{}

	// cek account sumber
	resultSourceAccount := c.CustomerRepo.GetCustomerSaldo(req.FromAccountNumber)
	if resultSourceAccount.Error != nil {
		if resultSourceAccount.Error == sql.ErrNoRows {
			result.Error = fmt.Errorf(helper.ErrorDataNotFound, "account sumber")
			return result
		}
		result.Error = resultSourceAccount.Error
		return result
	}

	// cek balance
	sourceAccount := resultSourceAccount.Result.(model.CustomerAccount)
	if sourceAccount.Balance < int64(req.Amount) {
		result.Error = fmt.Errorf(helper.ErrorAccountBalance)
		return result
	}

	// cek account tujuan
	resultDestinationAccount := c.CustomerRepo.GetCustomerSaldo(req.ToAccountNumber)
	if resultDestinationAccount.Error != nil {
		if resultDestinationAccount.Error == sql.ErrNoRows {
			result.Error = fmt.Errorf(helper.ErrorDataNotFound, "account tujuan")
			return result
		}
		result.Error = resultDestinationAccount.Error
		return result
	}

	resultTransfer := c.CustomerRepo.TransferSaldo(req.FromAccountNumber, req.ToAccountNumber, req.Amount)
	if resultTransfer.Error != nil {
		result.Error = resultTransfer.Error
		return result
	}

	return ResultUseCase{Result: nil}
}
