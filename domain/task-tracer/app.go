package tasktracer

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func App() {
	options := []string{"Add Task", "Update Task", "Delete Task", "List All Tasks", "List Done Tasks", "List In Progress Tasks", "List Todo Tasks", "Exit"}
	reader := bufio.NewReader(os.Stdin)
	CreateTaskFileIfNeeded()
	for {
		fmt.Println("\nSelect an option:")
		for i, opt := range options {
			if i == len(options)-1 {
				fmt.Printf("%d. %s\n", 0, opt)
				break
			}
			fmt.Printf("%d. %s\n", i+1, opt)
		}
		fmt.Print("Enter your choice: ")
		choice, _ := reader.ReadString('\n')
		choiceInt, _ := strconv.Atoi(strings.TrimSpace(choice))

		switch choiceInt {
		case 1:
			fmt.Println("Adding task...")
		case 2:
			fmt.Println("Updating task...")
		case 3:
			fmt.Println("Deleting task...")
		case 4:
			fmt.Println("Listing all tasks...")
		case 5:
			fmt.Println("Listing done tasks...")
		case 6:
			fmt.Println("Listing in progress tasks...")
		case 7:
			fmt.Println("Listing todo tasks...")
		case 0:
			fmt.Println("Exiting...")
			return
		}
	}
}
