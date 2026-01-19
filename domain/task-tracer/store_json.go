package tasktracer

import (
	"encoding/json"
	"fmt"
	"os"
)

const tasksFilePath = "tasks.json"

// loadTasks reads tasks from the JSON file in the current directory.
func loadTasks() []Task {
	content, err := os.ReadFile(tasksFilePath)
	if err != nil {
		// Keep current behavior (print error, return empty slice).
		fmt.Printf("Error reading file: %v\n", err)
		return []Task{}
	}

	// Handle empty file gracefully.
	if len(content) == 0 {
		return []Task{}
	}

	var tasks []Task
	if err := json.Unmarshal(content, &tasks); err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return []Task{}
	}

	return tasks
}

// saveTasks writes tasks to the JSON file in the current directory.
func saveTasks(tasks []Task) error {
	content, err := json.Marshal(tasks)
	if err != nil {
		return fmt.Errorf("marshalling JSON: %w", err)
	}
	return os.WriteFile(tasksFilePath, content, 0644)
}
