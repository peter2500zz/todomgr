package todofile

import (
	"fmt"
	"os"
	"path/filepath"
)

// GetTodoFilePath looks up the first todo file in the current directory and its parents.
//
// It returns the path to the file if found, or an error if not found.
func GetTodoFilePath() (string, error) {
	// get current working directory
	start, _ := os.Getwd()

	return lookUpTodoFile(start)
}

func lookUpTodoFile(dir string) (string, error) {
	todoFile := filepath.Join(dir, TODO_FILE_NAME)

	// find it
	if _, err := os.Stat(todoFile); err == nil {
		return todoFile, nil
	}

	parent := filepath.Dir(dir)

	// reached root
	if parent == dir {
		return "", fmt.Errorf("%s not found", TODO_FILE_NAME)
	}

	return lookUpTodoFile(parent)
}



