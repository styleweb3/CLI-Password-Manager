package cmd

import (
	"fmt"
	"time"

	"github.com/styleweb3/CLI-pwd-manager/models"
	"github.com/styleweb3/CLI-pwd-manager/internal/store"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add [service]",
	Short: "Add a new password entry",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		entry := models.PasswordEntry{}
		entry.ServiceName = args[0]
		entry.CreatedAt = time.Now()

		fmt.Print("Enter LogIn: ")
		fmt.Scanln(&entry.LogIn)

		fmt.Print("Enter password: ")
        fmt.Scanln(&entry.Password)

		entries, err := store.Load()
        if err != nil {
            fmt.Println("Error loading vault:", err)
            return
        }

        entries = append(entries, entry)

        err = store.Save(entries)
        if err != nil {
            fmt.Println("Error saving vault:", err)
            return
        }

        fmt.Printf("\n• Entry saved for: %s\n", entry.ServiceName)
	},
}

func init() {
	rootCmd.AddCommand(addCmd)
}
