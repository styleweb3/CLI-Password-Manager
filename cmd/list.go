package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/styleweb3/CLI-pwd-manager/internal/store"
)

var listCMD = &cobra.Command{
	Use:   "list",
	Short: "List all saved password entries",
	Run: func(cmd *cobra.Command, args []string) {
		entries, err := store.Load()
		if err != nil {
			fmt.Println("Error loading vault:", err)
			return
		}

		if len(entries) == 0 {
			fmt.Println("No entries found.")
			return
		}

		for _, entry := range entries {
			fmt.Println("-----------------------------")
			fmt.Printf("• Service:    %s", entry.ServiceName)
			fmt.Printf("\n• Log-In:      %s", entry.LogIn)
			fmt.Printf("\n• Created At: %s\n", entry.CreatedAt.Format("2006-01-02 15:04"))
		}
	
			fmt.Println("-----------------------------")

	},
}

func init() {
	rootCmd.AddCommand(listCMD)
}
