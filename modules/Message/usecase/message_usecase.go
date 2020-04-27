package usecase

import "WarungPintarAPI/modules/Message/repo"

type MessageUseCaseImpl struct {
	MessageRepo repo.MessageRepository
}

func NewMessageUseCase(messageRepo repo.MessageRepository) MessageUseCase {
	return &MessageUseCaseImpl{MessageRepo: messageRepo}
}

func (m *MessageUseCaseImpl) AddMessage(message string) ResultUseCase {

	result := m.MessageRepo.Add(message)
	if result.Error != nil {
		return ResultUseCase{Error: result.Error}
	}

	return ResultUseCase{Result: result.Result}
}

func (m *MessageUseCaseImpl) GetMessage() ResultUseCase {

	result := m.MessageRepo.Get()
	if result.Error != nil {
		return ResultUseCase{Error: result.Error}
	}

	return ResultUseCase{Result: result.Result}
}
