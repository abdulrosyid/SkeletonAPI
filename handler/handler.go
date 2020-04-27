package handler

import (
	"WarungPintarAPI/modules/Message/model"
	"WarungPintarAPI/modules/Message/presenter"
	"WarungPintarAPI/modules/Message/repo"
	"WarungPintarAPI/modules/Message/usecase"
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
