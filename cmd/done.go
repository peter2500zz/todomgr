package cmd

import (
	"fmt"
	"strconv"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var doneTodo = &cobra.Command{
	Use: "done",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Error: done need more than 1 argument\n")
		}

		path, err := todofile.GetTodoFilePath()
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		items, err := todofile.ReadTodo(path)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		for _, arg := range args {
			number, err := strconv.Atoi(arg)
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			if number <= 0 {
				fmt.Printf("Invalid number %d\n", number)
				return
			}
			if number > len(items) {
				fmt.Printf("You only have %d todo(s), but the given number is %d\n", len(items), number)
				return
			}
			if items[number - 1].Completed {
				fmt.Printf("Task \"%s\" is already completed\n", items[number - 1].Description)
				return
			}
			items[number - 1].Completed = true
		}

		todofile.WriteTodoFile(path, items)
	},
}
