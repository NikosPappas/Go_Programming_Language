package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string
	Done        bool
}

var tasks []Task

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		fmt.Println("To-Do Application")
		fmt.Println("1. Insert Task")
		fmt.Println("2. Mark Task as Done")
		fmt.Println("3. List Tasks")
		fmt.Println("4. Exit")
		fmt.Print("Choose an option: ")

		option, _ := reader.ReadString('\n')
		option = strings.TrimSpace(option)

		switch option {
		case "1":
			insertTask(reader)
		case "2":
			markTaskAsDone(reader)
		case "3":
			listTasks()
		case "4":
			fmt.Println("Exiting...")
			return
		default:
			fmt.Println("Invalid option. Please try again.")
		}
	}
}

func insertTask(reader *bufio.Reader) {
	fmt.Print("Enter task description: ")
	description, _ := reader.ReadString('\n')
	description = strings.TrimSpace(description)

	tasks = append(tasks, Task{Description: description, Done: false})
	fmt.Println("Task added!")
}

func markTaskAsDone(reader *bufio.Reader) {
	listTasks()
	fmt.Print("Enter the task number to mark as done: ")
	taskNumStr, _ := reader.ReadString('\n')
	taskNumStr = strings.TrimSpace(taskNumStr)

	var taskNum int
	_, err := fmt.Sscanf(taskNumStr, "%d", &taskNum)
	if err != nil || taskNum < 1 || taskNum > len(tasks) {
		fmt.Println("Invalid task number.")
		return
	}

	tasks[taskNum-1].Done = true
	fmt.Println("Task marked as done!")
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}

	fmt.Println("Tasks:")
	for i, task := range tasks {
		status := " "
		if task.Done {
			status = "✓"
		}
		fmt.Printf("%d. [%s] %s", i+1, status, task.Description)
	}
}

