package cmd

import (
	"github.com/hibiken/asynq"
	"github.com/spf13/cobra"
	"github.com/titanabrian/asynq-poc/client"
	"github.com/titanabrian/asynq-poc/message"
)

var (
	REDIS_SERVER = "localhost:6379"

	publishCmd = &cobra.Command{
		Use:   "enqueue",
		Short: "Queue message to broker",
		Run:   publishMessage,
	}
)

func init() {
	rootCmd.AddCommand(publishCmd)
}

func publishMessage(cmd *cobra.Command, args []string) {
	conn := asynq.NewClient(asynq.RedisClientOpt{
		Addr: REDIS_SERVER,
	})

	c := client.NewClient(*conn)

	type msg struct {
		ID        string
		Recepient string
	}

	m1 := message.NewMessage(
		"email:delivery",
		msg{
			ID:        "BAFA52A8-E444-4CF8-888C-93D68F7E912E",
			Recepient: "titan@xendit.co",
		})

	m2 := message.NewMessage(
		"sms:delivery",
		msg{
			ID:        "BAFA52A8-E444-4CF8-888C-93D68F7E912E",
			Recepient: "08xxxxxxx",
		},
	)

	c.Publish(m1)
	c.Publish(m2)

	defer conn.Close()
}
