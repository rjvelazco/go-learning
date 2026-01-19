package tasktracer

import (
	"errors"
	"os"
)

// File Utils (no Task knowledge)
func CreateTaskFileIfNeeded() {
	if fileExists(tasksFilePath) {
		return
	}
	// Keep prior behavior: best-effort create; ignore error for now.
	_, _ = os.Create(tasksFilePath)
}

// Checks if a file exists
func fileExists(filename string) bool {
	_, err := os.Stat(filename)
	if err == nil {
		return true
	}
	return !errors.Is(err, os.ErrNotExist)
}
