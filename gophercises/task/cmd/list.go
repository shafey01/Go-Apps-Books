package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCmd)
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List uncompeleted tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		db, err := setupDB()

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		tasks, err := ListTasks(db, false)
		if err != nil {
			exitf("%v\n", err)

		}
		if len(tasks) == 0 {
			fmt.Println("You don't have any incompleted tasks.")
			os.Exit(0)
		}
		fmt.Println("Your Tasks: ")
		for _, t := range tasks {
			fmt.Printf("%d. %s\n", t.ID, t.Details)
		}
	},
}
