package cmd

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/styleweb3/CLI-pwd-manager/internal/store"
	"github.com/styleweb3/CLI-pwd-manager/models"
)

var deleteCMD = &cobra.Command{
	Use: 	"del [service]",
	Short: 	"Delete a password entry by service name",
	Args: 	cobra.ExactArgs(1),
	Run: 	func(cmd *cobra.Command, args []string){
				searchTerm := strings.ToLower(args[0])

        		entries, err := store.Load()
        		
				if err != nil {
            		fmt.Println("Error loading vault:", err)
            		return
      			}

				var filtered []models.PasswordEntry
				found := false
				
				for _, entry := range entries{
					if strings.ToLower(entry.ServiceName) == searchTerm {
                		found = true
                		continue
            		}
            		filtered = append(filtered, entry)
				}

				if !found {
					fmt.Printf("No entry found for: %s\n", args[0])
					return
				}

				err = store.Save(filtered)
        		if err != nil {
          			fmt.Println("Error saving vault:", err)
            		return
        		}

       			fmt.Printf("• Entry deleted for: %s\n", args[0])
			},
}

func init(){
	rootCmd.AddCommand(deleteCMD)
}