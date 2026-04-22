package cmd

import (
	"fmt"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var newTodo = &cobra.Command{
	Use: "new \"<something todo>\"",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 1 {
			fmt.Printf("Usage: todo new \"<something todo>\"\n")
			return
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

		items = append(items, todofile.TodoItem{
			Description: args[0],
			Completed:   false,
		})

		// t := time.Now()

		// fmt.Printf("Add new todo item %s at %s\n", args[0], t.Format(time.RFC3339))

		err = todofile.WriteTodoFile(path, items)
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		fmt.Printf("Added new todo \"%s\"\n", args[0])
		fmt.Printf("Now you have %d todo(s) to be completed.\n", items.CountIncomplete())
	},
}
