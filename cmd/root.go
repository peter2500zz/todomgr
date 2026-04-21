package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "todo",
	Run: func(cmd *cobra.Command, args []string) {
		// print the next suggested todo
		fmt.Println("Next:\nCook my dinner.\n\nFinished 3/5 tasks today. Keep it up!")
	},
}

func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(listTodo)
}
