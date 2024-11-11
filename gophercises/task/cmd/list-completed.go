package cmd

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(listCompletedCmd)
}

var listCompletedCmd = &cobra.Command{
	Use:   "done",
	Short: "List compeleted tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		db, err := setupDB()

		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		tasks, err := ListTasks(db, true)
		if err != nil {
			exitf("%v\n", err)

		}
		if len(tasks) == 0 {
			fmt.Println("You don't have any completed tasks.")
			os.Exit(0)
		}
		fmt.Println("Your Completed Tasks: ")
		for _, t := range tasks {
			fmt.Printf("%d. %s\n", t.ID, t.Details)
		}
	},
}
