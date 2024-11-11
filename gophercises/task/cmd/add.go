package cmd

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(addCmd)
}

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "Add tasks",
	Run: func(cmd *cobra.Command, args []string) {
		// Do Stuff Here
		db, err := setupDB()
		if err != nil {
			log.Fatal(err)
		}
		defer db.Close()

		task := Task{Details: strings.Join(args, " ")}
		if len(args) == 0 {
			fmt.Printf("Please write task to add!")
			os.Exit(1)
		}
		if err := SetTask(db, &task); err != nil {
			exitf("%v\n", err)
		}

	},
}
