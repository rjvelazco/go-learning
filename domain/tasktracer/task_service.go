package tasktracer

import (
	"fmt"
	"strconv"
)

// AddTask appends a new task to the current task list and persists it to disk.
func AddTask(task Task) {
	tasks := loadTasks()
	tasks = append(tasks, task)
	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}

// UpdateTask updates the title and description of the task with the given id,
// then persists the updated task list to disk.
func UpdateTask(id int, title string, description string) {
	tasks := loadTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			if description != "" {
				tasks[i].Description = description
			}
			break
		}
	}
	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}

// DeleteTask removes the task with the given id from the task list and persists
// the updated list to disk.
func DeleteTask(id int) {
	tasks := loadTasks()
	newTasks := []Task{}
	for i, task := range tasks {
		if task.ID == id {
			newTasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	if err := saveTasks(newTasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}

// updateTaskStatus updates the status of the task with the given id,
// then persists the updated task list to disk.
func updateTaskStatus(id string, status string) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: Invalid ID\n")
		fmt.Printf("Usage: task-cli mark-in-progress <id>\n")
		return
	}

	tasks := loadTasks()
	if status != "done" && status != "in-progress" {
		fmt.Printf("Invalid status: %s\n", status)
		fmt.Printf("Usage: task-cli mark-in-progress <id>\n")
		return
	}

	for i, task := range tasks {
		if task.ID == idInt {
			tasks[i].Status = status
			break
		}
	}

	if err := saveTasks(tasks); err != nil {
		fmt.Printf("Error saving tasks: %v\n", err)
	}
}
