package tasktracer

import (
	"errors"
	"os"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Status      string
}

func CreateTask(id int, title string, description string) *Task {
	return &Task{
		ID:          id,
		Title:       title,
		Description: description,
		Status:      "new",
	}
}

func (t *Task) UpdateStatus(status string) {
	t.Status = status
}

func (t *Task) UpdateTitle(title string) {
	t.Title = title
}

func (t *Task) UpdateDescription(description string) {
	t.Description = description
}

func CreateTaskFileIfNeeded() {
	if fileExists("tasks.json") {
		return
	}

	os.Create("tasks.json")
}

func fileExists(filename string) bool {
	_, err := os.Stat(filename)

	if err == nil {
		return true
	}

	if errors.Is(err, os.ErrNotExist) {
		return false
	}

	return false
}
