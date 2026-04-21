package cmd

import (
	"fmt"
	"todomgr/internal/todofile"

	"github.com/spf13/cobra"
)

var listTodo = &cobra.Command{
	Use: "list",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(todofile.GetTodoFilePath())
	},
}