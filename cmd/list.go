package cmd

import (
	"fmt"
	"todomgr/internal/todofile"

	"charm.land/lipgloss/v2"
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
			var descStyle = lipgloss.NewStyle()
			var checkmark string = " "

			if item.Completed {
				// weak and strikethrough
				descStyle = descStyle.Faint(true)
				descStyle = descStyle.StrikethroughSpaces(true)

				// green checkmark
				checkmark = lipgloss.
					NewStyle().
					Foreground(lipgloss.Color("#00FF00")).
					Render("x")
			}

			fmt.Printf(
				"%d. [%s] %s\n",
				i+1,
				checkmark,
				descStyle.Render(item.Description),
			)
		}
	},
}
