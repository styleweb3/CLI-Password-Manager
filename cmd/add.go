package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [service]",
	Short: "Add a new password entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Adding entry for: %s\n", args[0])
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
