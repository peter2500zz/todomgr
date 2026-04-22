package cmd

import (
	"fmt"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var listTodo = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
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

		for i, item := range items {

			fmt.Printf(
				"%d. %s\n",
				i+1,
				item.StringPetty(),
			)
		}
	},
}
