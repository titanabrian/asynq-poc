package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "template",
		Short: "Template is tempalte management application",
	}
)

func Execute() {
	fmt.Printf("cobra command started\n")
	if err := rootCmd.Execute(); err != nil {
		panic(err)
	}
}
