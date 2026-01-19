package tasktracer

import (
	"crypto/rand"
	"fmt"
	"math/big"
	"strconv"
)

// addNewTask adds a new task to the task list and persists it to disk.
func AddNewTask(taskData []string) {
	if len(taskData) < 2 {
		fmt.Printf("Error: Missing title and description\n")
		fmt.Printf("Usage: task-cli add \"title\" \"description\"\n")
		return
	}

	title := taskData[0]
	description := taskData[1]
	randomID, _ := rand.Int(rand.Reader, big.NewInt(1000000))

	AddTask(Task{
		ID:          int(randomID.Int64()),
		Title:       title,
		Description: description,
		Status:      "todo",
	})
}

// handleUpdate updates the title and description of the task with the given id,
// then persists the updated task list to disk.
func UpdateTaskById(id string, title string, description string) {
	idInt, err := strconv.Atoi(id)
	if err != nil {
		fmt.Printf("Error: Invalid ID\n")
		fmt.Printf("Usage: task-cli update <id> \"title\" \"description\"\n")
		return
	}
	if idInt < 0 || title == "" || description == "" {
		fmt.Printf("Error: Invalid ID, title or description\n")
		fmt.Printf("Usage: task-cli update <id> \"title\" \"description\"\n")
		return
	}

	UpdateTask(idInt, title, description)
}

// handleDelete deletes the task with the given id from the task list and persists
// the updated list to disk.
func handleDelete(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting ID to int: %v\n", err)
		return
	}
	DeleteTask(id)
}

// ListAllTasks lists all tasks in the task list.
func ListAllTasks() {
	tasks := loadTasks()
	fmt.Println("Listing all tasks:")
	fmt.Println("--------------------------------")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
	fmt.Println("--------------------------------")
}

// ListByStatusTasks lists all tasks in the task list with the given status.
func ListByStatusTasks(status string) {
	if status != "done" && status != "todo" && status != "in-progress" {
		fmt.Printf("Invalid status: %s\n", status)
		fmt.Printf("Usage: task-cli list [done|todo|in-progress]\n")
		return
	}

	isEmpty := true
	tasks := loadTasks()
	fmt.Printf("Listing %s tasks:\n", status)
	fmt.Println("--------------------------------")
	for _, task := range tasks {
		if task.Status == status {
			isEmpty = false
			fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
		}
	}

	if isEmpty {
		fmt.Printf("No tasks found in: \"%s\" Status\n", status)
	}
	fmt.Println("--------------------------------")
}
