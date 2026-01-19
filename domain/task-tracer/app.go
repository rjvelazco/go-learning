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
		handleAdd(rest)
	case "update":
		// handleUpdate(rest)
	case "delete":
		// handleDelete(rest)
	case "list":
		fmt.Println("list")
	case "mark-in-progress":
		fmt.Println("mark-in-progress")
	case "mark-done":
		fmt.Println("mark-done")
	default:
		fmt.Fprintf(os.Stderr, "Invalid command: %s\n\n", cmd)
		usage()
		os.Exit(1)
	}
}

func usage() {
	fmt.Fprintln(os.Stderr, "Usage:")
	fmt.Fprintln(os.Stderr, "  task-cli add \"description\"")
	fmt.Fprintln(os.Stderr, "  task-cli update <id> \"description\"")
	fmt.Fprintln(os.Stderr, "  task-cli delete <id>")
	fmt.Fprintln(os.Stderr, "  task-cli mark-in-progress <id>")
	fmt.Fprintln(os.Stderr, "  task-cli mark-done <id>")
	fmt.Fprintln(os.Stderr, "  task-cli list [done|todo|in-progress]")
}
