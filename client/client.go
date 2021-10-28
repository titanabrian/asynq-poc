package client

import (
	"log"

	"github.com/hibiken/asynq"
	"github.com/titanabrian/asynq-poc/message"
)

type Client struct {
	broker asynq.Client
}

func NewClient(broker asynq.Client) *Client {
	return &Client{broker: broker}
}

func (c *Client) Publish(message message.IMessage) error {
	task, err := message.NewTask()
	if err != nil {
		return err
	}

	info, err := c.broker.Enqueue(task)
	if err != nil {
		return err
	}

	log.Printf("enqueued task: id=%s queue=%s", info.ID, info.Queue)

	return nil
}
