package message

import (
	"encoding/json"

	"github.com/hibiken/asynq"
)

type IMessage interface {
	NewTask() (*asynq.Task, error)
}

type MessageImpl struct {
	Name    string
	Message interface{}
}

func NewMessage(name string, msg interface{}) *MessageImpl {
	return &MessageImpl{
		Name:    name,
		Message: msg,
	}
}

func (m *MessageImpl) NewTask() (*asynq.Task, error) {
	payload, err := json.Marshal(m.Message)
	if err != nil {
		return nil, err
	}

	return asynq.NewTask(m.Name, payload), nil
}
