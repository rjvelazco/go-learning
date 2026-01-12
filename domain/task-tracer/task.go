package tasktracer

import (
	"bufio"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"
	"os"
	"strings"
)

type Task struct {
	ID          int    `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Status      string `json:"status"`
}

func CreateTask() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter task title: ")
	title, _ := reader.ReadString('\n')
	title = strings.TrimSpace(title)
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)
	randomID, _ := rand.Int(rand.Reader, big.NewInt(1000000))

	AddTask(Task{
		ID:          int(randomID.Int64()),
		Title:       title,
		Description: description,
		Status:      "new",
	})
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

func ListTasks() {
	tasks := getTasks()
	fmt.Println("Listing all tasks:")
	fmt.Println("--------------------------------")
	for _, task := range tasks {
		fmt.Printf("ID: %d, Title: %s, Description: %s, Status: %s\n", task.ID, task.Title, task.Description, task.Status)
	}
	fmt.Println("--------------------------------")
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
