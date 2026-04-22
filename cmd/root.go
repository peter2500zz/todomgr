package cmd

import (
	"fmt"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "todo",
	Run: func(cmd *cobra.Command, args []string) {
		// print the next suggested todo
		path, err := todofile.GetTodoFilePath()
		if err != nil {
			fmt.Printf("Type todo help for usage information.\n")
			return
		}

		items, err := todofile.ReadTodo(path)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		var firstIncomplete *todofile.TodoItem
		for _, item := range items {
			if !item.Completed {
				firstIncomplete = &item
				break
			}
		}
		if firstIncomplete != nil {
			fmt.Printf("Next:\n%s\n", firstIncomplete.Description)
		} else {
			fmt.Println("No more todos! Great work!")
		}

	},
}

func Execute() {
	_ = rootCmd.Execute()
}

func init() {
	rootCmd.AddCommand(listTodo)
	rootCmd.AddCommand(doneTodo)
	rootCmd.AddCommand(undoTodo)
	rootCmd.AddCommand(newTodo)
	rootCmd.AddCommand(swapTodo)
	rootCmd.AddCommand(delTodo)
}
