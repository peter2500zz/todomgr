package todofile

import (
	"fmt"
	"strings"

	"charm.land/lipgloss/v2"
)

const TODO_FILE_NAME = "todo.md"

type TodoItem struct {
	Description string
	Completed bool
}

// String returns a string representation of the todo item, like "[ ] Something todo"
func (t TodoItem) String() string {
	// "[ ] Something todo"

	completeMark := "[ ]"
	if t.Completed {
		completeMark = "[x]"
	}

	return fmt.Sprintf("%s %s", completeMark, t.Description)
}

// StringPetty returns a string that looks pretty in terminal
func (t TodoItem) StringPetty() string {
	// string with color

	var stringPetty strings.Builder = strings.Builder{}
	stringPetty.WriteString("[")
	if t.Completed {
		stringPetty.WriteString(lipgloss.NewStyle().Foreground(lipgloss.Color("#00FF00")).Render("x"))
	} else {
		stringPetty.WriteString(" ")
	}
	stringPetty.WriteString("] ")

	if !t.Completed {
		stringPetty.WriteString(t.Description)
	} else {
		stringPetty.WriteString(lipgloss.NewStyle().Faint(true).StrikethroughSpaces(true).Render(t.Description))
	}

	return stringPetty.String()
}

type TodoList []TodoItem

// FirstIncomplete returns the first incomplete todo item, or nil if all items are completed.
func (t TodoList) FirstIncomplete() *TodoItem {
	for _, item := range t {
		if !item.Completed {
			return &item
		}
	}
	return nil
}

func (t TodoList) CountIncomplete() int {
	count := 0
	for _, item := range t {
		if !item.Completed {
			count++
		}
	}
	return count
}

func (t TodoList) CountCompleted() int {
	count := 0
	for _, item := range t {
		if item.Completed {
			count++
		}
	}
	return count
}
