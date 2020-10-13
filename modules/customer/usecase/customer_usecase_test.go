package usecase

import (
	"SkeletonAPI/modules/customer/model"
	"SkeletonAPI/modules/customer/repo"
	repoMock "SkeletonAPI/modules/customer/repo/mocks"
	"database/sql"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
)

func TestNewCustomerUseCase(t *testing.T) {
	mockRepo := new(repoMock.CustomerRepository)
	_ = NewCustomerUseCase(mockRepo)
	assert.NoError(t, nil)
}

func TestCustomerUseCaseImpl_CheckSaldo(t *testing.T) {
	t.Run("Test Success CheckSaldo", func(t *testing.T) {
		customerAccount := model.CustomerAccount{}

		resultRepo := repo.ResultRepository{}
		resultRepo.Result = customerAccount

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.CheckSaldo(1)
		assert.NoError(t, result.Error)
	})

	t.Run("Test Failed CheckSaldo ErrConnDone", func(t *testing.T) {
		resultRepo := repo.ResultRepository{}
		resultRepo.Error = sql.ErrConnDone

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.CheckSaldo(1)
		assert.Error(t, result.Error)
	})

	t.Run("Test Failed CheckSaldo ErrNoRows", func(t *testing.T) {
		resultRepo := repo.ResultRepository{}
		resultRepo.Error = sql.ErrNoRows

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.CheckSaldo(1)
		assert.Error(t, result.Error)
	})
}

func TestCustomerUseCaseImpl_TransferSaldo(t *testing.T) {
	t.Run("Test Success TransferSaldo", func(t *testing.T) {
		customerAccount := model.CustomerAccount{}
		req := model.TransferRequest{}

		resultRepo := repo.ResultRepository{}
		resultRepo.Result = customerAccount

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Return(resultRepo)
		mockRepo.On("TransferSaldo", mock.Anything, mock.Anything, mock.Anything).Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.TransferSaldo(req)
		assert.NoError(t, result.Error)
	})

	t.Run("Test Failed CheckSaldo cek account sumber ErrNoRows", func(t *testing.T) {
		resultRepo := repo.ResultRepository{}
		req := model.TransferRequest{}
		resultRepo.Error = sql.ErrNoRows

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.TransferSaldo(req)
		assert.Error(t, result.Error)
	})

	t.Run("Test Failed CheckSaldo cek account sumber ErrConnDone", func(t *testing.T) {
		resultRepo := repo.ResultRepository{}
		req := model.TransferRequest{}
		resultRepo.Error = sql.ErrConnDone

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.TransferSaldo(req)
		assert.Error(t, result.Error)
	})

	t.Run("Test Failed TransferSaldo Balance", func(t *testing.T) {
		customerAccount := model.CustomerAccount{}
		req := model.TransferRequest{}

		resultRepo := repo.ResultRepository{}
		resultRepo.Result = customerAccount

		resultRepoError := repo.ResultRepository{}
		resultRepoError.Error = sql.ErrNoRows

		mockRepo := new(repoMock.CustomerRepository)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepo)
		mockRepo.On("GetCustomerSaldo", mock.Anything).Once().Return(resultRepoError)
		mockRepo.On("TransferSaldo", mock.Anything, mock.Anything, mock.Anything).Return(resultRepo)

		customerUseCaseImpl := new(CustomerUseCaseImpl)
		customerUseCaseImpl.CustomerRepo = mockRepo

		result := customerUseCaseImpl.TransferSaldo(req)
		assert.NoError(t, result.Error)
	})
}
