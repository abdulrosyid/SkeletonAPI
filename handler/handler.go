package handler

import (
	"SkeletonAPI/config"
	"SkeletonAPI/modules/customer/presenter"
	"SkeletonAPI/modules/customer/repo"
	"SkeletonAPI/modules/customer/usecase"
)

type Service struct {
	CustomerHandler *presenter.HTTPMessageHandler
	CustomerUseCase usecase.CustomerUseCase
}

func MakeHandler() *Service {
	// ini connection DB
	writeMysqlDB := config.PostgreSqlDB()

	customerRepo := repo.NewCustomerRepo(writeMysqlDB)
	customerUseCase := usecase.NewCustomerUseCase(customerRepo)
	customerHandler := presenter.NewHTTPHandler(customerUseCase)

	return &Service{
		CustomerHandler: customerHandler,
		CustomerUseCase: customerUseCase,
	}
}
