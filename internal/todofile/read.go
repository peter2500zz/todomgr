package todofile

import (
	"os"
	"strings"
)

func ReadTodo(path string) (TodoList, error) {
	content, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(content), "\n")

	var items TodoList
	for _, line := range lines {
		// skip empty lines
		if line == "" {
			continue
		}

		runes := []rune(line)

		if len(runes) < 5 {
			// what's wrong with this line?
			continue
		}

		prefix := string(runes[0:5])
		desc := strings.TrimSpace(string(runes[5:]))

		switch prefix {
		case "- [ ]":
			items = append(items, TodoItem{
				Description: desc,
				Completed:   false,
			})
		case "- [x]":
			items = append(items, TodoItem{
				Description: desc,
				Completed:   true,
			})
		default:
			// weird line, skip
			continue
		}
	}

	return items, nil
}
