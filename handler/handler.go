package handler

import (
	"SkeletonAPI/modules/Message/model"
	"SkeletonAPI/modules/Message/presenter"
	"SkeletonAPI/modules/Message/repo"
	"SkeletonAPI/modules/Message/usecase"
)

type Service struct {
	MessageHandler *presenter.HTTPMessageHandler
	MessageUseCase usecase.MessageUseCase
}

func MakeHandler() *Service {
	message := model.ListMessage{}
	messageRepo := repo.NewMessageDataRepo(&message)
	messageUseCase := usecase.NewMessageUseCase(messageRepo)
	messageHandler := presenter.NewHTTPHandler(messageUseCase)

	return &Service{
		MessageHandler: messageHandler,
		MessageUseCase: messageUseCase,
	}
}
