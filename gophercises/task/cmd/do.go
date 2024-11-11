package cmd

import (
	"fmt"
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(doCmd)
}

var doCmd = &cobra.Command{
	Use:   "do",
	Short: "Mark task as compeleted.",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		db, err := setupDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		if len(args) == 0 {
			fmt.Printf("Please write task to add!")
			os.Exit(1)

		}
		id, err := strconv.Atoi(args[0])
		if err != nil {

			exitf("%v\n", err)
		}
		task := Task{ID: id}
		if err := MarkTaskAsCompleted(db, &task); err != nil {
			exitf("%v\n", err)
		}
	},
}
