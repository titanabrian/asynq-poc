package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/titanabrian/asynq-poc/worker"
)

var (
	workerCmd = &cobra.Command{
		Use:   "worker",
		Short: "Star worker",
		Run:   processMessage,
	}
)

func init() {
	rootCmd.AddCommand(workerCmd)
}

func processMessage(cmd *cobra.Command, args []string) {
	w := worker.NewWorker(REDIS_SERVER, "email:delivery")
	if err := w.Start(); err != nil {
		fmt.Printf("could not run server: %v", err)
	}
}
