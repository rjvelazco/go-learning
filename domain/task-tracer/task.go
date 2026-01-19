package tasktracer

import (
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strconv"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func handleAdd(taskData []string) {
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
		Status:      "new",
	})
}

func handleUpdate(args []string) {
	if len(args) < 3 {
		fmt.Printf("Error: Missing ID, title and description\n")
		fmt.Printf("Usage: task-cli update <id> \"title\" \"description\"\n")
		return
	}

	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting ID to int: %v\n", err)
		return
	}

	title := args[1]
	description := args[2]

	UpdateTask(int(id), title, description)
}

func handleDelete(args []string) {
	id, err := strconv.Atoi(args[0])
	if err != nil {
		fmt.Printf("Error converting ID to int: %v\n", err)
		return
	}
	DeleteTask(int(id))
}

func listTasks(status string) {
	tasks := getTasks()
	fmt.Println("Listing all tasks:")
	fmt.Println("--------------------------------")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
	fmt.Println("--------------------------------")
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

func ListAllTasks() {
	tasks := getTasks()
	fmt.Println("Listing all tasks:")
	fmt.Println("--------------------------------")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
	fmt.Println("--------------------------------")
}

func ListByStatusTasks(status string) {
	if status != "done" && status != "todo" && status != "in-progress" {
		fmt.Printf("Invalid status: %s\n", status)
		fmt.Printf("Usage: task-cli list [done|todo|in-progress]\n")
		return
	}

	tasks := getTasks()
	fmt.Printf("Listing %s tasks:\n", status)
	fmt.Println("--------------------------------")
	for _, task := range tasks {
		if task.Status == status {
			fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
		}
	}
	fmt.Println("--------------------------------")
}

// Task Utils
func AddTask(task Task) {
	tasks := getTasks()
	tasks = append(tasks, task)
	content, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}
	os.WriteFile("tasks.json", content, 0644)
}

func UpdateTask(id int, title string, description string) {
	tasks := getTasks()
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Title = title
			tasks[i].Description = description
			break
		}
	}
	content, err := json.Marshal(tasks)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}
	os.WriteFile("tasks.json", content, 0644)
}

func DeleteTask(id int) {
	tasks := getTasks()
	newTasks := []Task{}
	for i, task := range tasks {
		if task.ID == id {
			newTasks = append(tasks[:i], tasks[i+1:]...)
			break
		}
	}
	content, err := json.Marshal(newTasks)
	if err != nil {
		fmt.Printf("Error marshalling JSON: %v\n", err)
		return
	}
	os.WriteFile("tasks.json", content, 0644)
}

func getTasks() []Task {
	filePath := "tasks.json"
	content, err := os.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error reading file: %v\n", err)
	}

	var tasks []Task
	err = json.Unmarshal(content, &tasks)
	if err != nil {
		fmt.Printf("Error unmarshalling JSON: %v\n", err)
		return []Task{}
	}

	return tasks
}

// File Utils
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
