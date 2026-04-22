package cmd

import (
	"fmt"
	"strconv"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var swapTodo = &cobra.Command{
	Use: "swap <index1> <index2>",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			fmt.Printf("Usage: todo swap <index1> <index2>\n")
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

		index1, err := strconv.Atoi(args[0])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}
		index2, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Printf("Error: %s\n", err)
			return
		}

		if index1 <= 0 || index2 <= 0 {
			fmt.Printf("Invalid number %d or %d\n", index1, index2)
			return
		}
		if index1 > len(items) || index2 > len(items) {
			fmt.Printf("You only have %d todo(s), but the given numbers are %d and %d\n", len(items), index1, index2)
			return
		}

		// swap them
		items[index1-1], items[index2-1] = items[index2-1], items[index1-1]
		todofile.WriteTodoFile(path, items)

		fmt.Printf("Tasks swapped successfully.\n")
	},
}
