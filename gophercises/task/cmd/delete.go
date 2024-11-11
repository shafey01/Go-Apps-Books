package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(deleteCmd)
}

var deleteCmd = &cobra.Command{
	Use:   "rm",
	Short: "Delete task by it's ID.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		db, err := setupDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if len(args) == 0 {
			fmt.Printf("Please write task ID to delete!")
			os.Exit(1)

		}
		id, err := strconv.Atoi(args[0])
		if err != nil {

			exitf("%v\n", err)
		}
		task := Task{ID: id}
		if err := Delete(db, &task); err != nil {
			exitf("%v\n", err)
		}
	},
}
