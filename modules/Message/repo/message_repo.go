package repo

import (
	"SkeletonAPI/modules/Message/model"
	"errors"
)

type MessageDataRepo struct {
	MessageData *model.ListMessage
}

func NewMessageDataRepo(messageData *model.ListMessage) *MessageDataRepo {
	return &MessageDataRepo{MessageData: messageData}
}

func (m *MessageDataRepo) Add(message string) ResultRepository {
	m.MessageData.Messages = append(m.MessageData.Messages, message)

	return ResultRepository{Result: m.MessageData.Messages}
}

func (m *MessageDataRepo) Get() ResultRepository {
	output := ResultRepository{}

	if len(m.MessageData.Messages) < 1 {
		err := errors.New("no data messages")
		output = ResultRepository{Error: err}
		return output
	}

	return ResultRepository{Result: m.MessageData.Messages}
}
