package worker

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/hibiken/asynq"
)

type EmailDeliveryPayload struct {
	ID      string
	Message interface{}
}

type Worker struct {
	RedisAddr   string
	MessageName string
}

func NewWorker(addr string, messageName string) *Worker {
	return &Worker{
		RedisAddr:   addr,
		MessageName: messageName,
	}
}

func (w *Worker) process(ctx context.Context, t *asynq.Task) error {
	var m EmailDeliveryPayload
	if err := json.Unmarshal(t.Payload(), &m); err != nil {
		return fmt.Errorf("json.Unmarshal failed: %v: %w", err, asynq.SkipRetry)
	}

	fmt.Printf("message ID: %s", m.ID)
	return nil
}

func (w *Worker) Start() error {
	srv := asynq.NewServer(asynq.RedisClientOpt{
		Addr: w.RedisAddr,
	}, asynq.Config{
		Queues: map[string]int{
			"default": 10,
		},
	})

	mux := asynq.NewServeMux()
	mux.HandleFunc(w.MessageName, w.process)

	if err := srv.Run(mux); err != nil {
		return err
	}

	return nil
}
