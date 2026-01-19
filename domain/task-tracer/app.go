package tasktracer

import (
	"fmt"
	"os"
)

func App() {
	CreateTaskFileIfNeeded()
	args := os.Args[1:]
	if len(args) == 0 {
		usage()
		os.Exit(1)
	}

	cmd := args[0]
	rest := args[1:]

	switch cmd {
	case "add":
		AddNewTask(rest)
	case "update":
		UpdateTaskById(rest[0], rest[1], rest[2])
	case "delete":
		handleDelete(rest)
	case "list":
		if len(rest) == 0 {
			ListAllTasks()
		} else {
			ListByStatusTasks(rest[0])
		}
	case "mark-in-progress":
		updateTaskStatus(rest[0], "in-progress")
	case "mark-done":
		updateTaskStatus(rest[0], "done")

	default:
		fmt.Fprintf(os.Stderr, "Invalid command: %s\n\n", cmd)
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  task-cli add \"title\" \"description\"")
	fmt.Fprintln(os.Stderr, "  task-cli update <id> \"title\" \"description\"")
	fmt.Fprintln(os.Stderr, "  task-cli delete <id>")
	fmt.Fprintln(os.Stderr, "  task-cli mark-in-progress <id>")
	fmt.Fprintln(os.Stderr, "  task-cli mark-done <id>")
	fmt.Fprintln(os.Stderr, "  task-cli list [done|todo|in-progress]")
}
