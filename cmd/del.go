package cmd

import (
	"fmt"
	"slices"
	"strconv"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var delTodo = &cobra.Command{
	Use: "del <index1> <index2> ...",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) < 1 {
			fmt.Printf("Error: del need more than 1 argument\n")
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
		}

		var newItems []todofile.TodoItem
		var deletedItems []todofile.TodoItem
		for i, item := range items {
			if !slices.Contains(args, strconv.Itoa(i + 1)) {
				newItems = append(newItems, item)
			} else {
				deletedItems = append(deletedItems, item)
			}
		}
		fmt.Printf("Deleted:\n")
		for _, item := range deletedItems {
			fmt.Printf("%s\n", item.Description)
		}

		todofile.WriteTodoFile(path, newItems)
	},
}

