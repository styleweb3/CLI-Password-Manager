package cmd

import (
	"fmt"
	"strings"

	"github.com/styleweb3/CLI-pwd-manager/internal/store"
	"github.com/spf13/cobra"
)

var getCMD = &cobra.Command{
	Use: 	"get [service]",
	Short: 	"Get a password entry by service name",
	Args: 	cobra.ExactArgs(1),
	Run: 	func(cmd *cobra.Command, args []string){
		searchTerm := strings.ToLower(args[0])

		entries, err := store.Load()
        if err != nil {
            fmt.Println("Error loading vault:", err)
            return
        }

		for _, entry := range entries {
            if strings.ToLower(entry.ServiceName) == searchTerm {
                fmt.Println("-----------------------------")
                fmt.Printf("• Service:    %s\n", entry.ServiceName)
                fmt.Printf("• Log-In:     %s\n", entry.LogIn)
                fmt.Printf("• Password:   %s\n", entry.Password)
                fmt.Printf("• Created At: %s\n", entry.CreatedAt.Format("2006-01-02 15:04"))
                fmt.Println("-----------------------------")
                return
            }
        }

		fmt.Printf("No entry found for: %s\n", args[0])

	} ,
}

func init() {
	rootCmd.AddCommand(getCMD)
}