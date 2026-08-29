package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/styleweb3/CLI-pwd-manager/models"
)

var addCmd = &cobra.Command{
	Use:   "add [service]",
	Short: "Add a new password entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entry := models.PasswordEntry{}

		fmt.Print("Enter LogIn: ")
		fmt.Scanln(&entry.LogIn)

		fmt.Print("Enter password: ")
        fmt.Scanln(&entry.Password)

		fmt.Printf("\nEntry saved for: %s\n", entry.ServiceName)
		fmt.Printf("• LogIn: %s", entry.LogIn)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
