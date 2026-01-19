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
	case "help", "-h", "--help":
		usage()
		return
	case "add":
		AddNewTask(rest)
	case "update":
		if len(rest) < 2 {
			fmt.Printf("Error: Missing ID and title\n")
			usage()
			os.Exit(1)
		}
		if len(rest) >= 3 {
			UpdateTaskById(rest[0], rest[1], rest[2])
		} else {
			// Support a shorter form: update <id> "title"
			UpdateTaskById(rest[0], rest[1], "")
		}
	case "delete":
		if len(rest) < 1 {
			fmt.Printf("Error: Missing ID\n")
			usage()
			os.Exit(1)
		}
		handleDelete(rest)
	case "list":
		if len(rest) == 0 {
			ListAllTasks()
		} else {
			ListByStatusTasks(rest[0])
		}
	case "mark-in-progress":
		if len(rest) < 1 {
			fmt.Printf("Error: Missing ID\n")
			usage()
			os.Exit(1)
		}
		updateTaskStatus(rest[0], "in-progress")
	case "mark-done":
		if len(rest) < 1 {
			fmt.Printf("Error: Missing ID\n")
			usage()
			os.Exit(1)
		}
		updateTaskStatus(rest[0], "done")

	default:
		fmt.Fprintf(os.Stderr, "Invalid command: %s\n\n", cmd)
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  task-cli add \"title\" [\"description\"]")
	fmt.Fprintln(os.Stderr, "  task-cli update <id> \"title\" [\"description\"]")
	fmt.Fprintln(os.Stderr, "  task-cli delete <id>")
	fmt.Fprintln(os.Stderr, "  task-cli mark-in-progress <id>")
	fmt.Fprintln(os.Stderr, "  task-cli mark-done <id>")
	fmt.Fprintln(os.Stderr, "  task-cli list [done|todo|in-progress]")
}
