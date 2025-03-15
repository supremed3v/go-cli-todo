package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type Task struct {
	Description string `json:"description"`
	Completed   bool   `json:"isCompleted"`
}

func main() {
	tasks := loadTasks()

	for {
		fmt.Println("\nTo-Do List")
		fmt.Println("1. Add task")
		fmt.Println("2. List task")
		fmt.Println("3. Mark Task as Completed")
		fmt.Println("4. Remove Task")
		fmt.Println("5. Exit")

		// Get user choice
		fmt.Print("Enter your choice: ")
		reader := bufio.NewReader(os.Stdin)
		choice, _ := reader.ReadString('\n')

		choice = strings.TrimSpace(choice)

		switch choice {
		case "1":
			addTask(&tasks)
		case "2":
			listTasks(tasks)
		case "3":
			markTaskAsCompleted(&tasks)
		case "4":
			removeTask(&tasks)
		case "5":
			saveTasks(tasks)
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice. Please try again.")
		}
	}
}

func loadTasks() []Task {
	file, err := os.ReadFile("tasks.json")
	if err != nil {
		return []Task{}
	}

	var tasks []Task
	// Unmarshal the json into a slice of tasks
	err = json.Unmarshal(file, &tasks)
	if err != nil {
		fmt.Println("error reading tasks: ", err)
	}

	return tasks
}

func saveTasks(tasks []Task) {
	data, err := json.MarshalIndent(tasks, "", " ")
	if err != nil {
		fmt.Println("Error Saving tasks:", err)
		return
	}

	// Write the data to file using os.writefile
	err = os.WriteFile("tasks.json", data, 0644)
	if err != nil {
		fmt.Println("Error writing tasks to file: ", err)
	}
}

func addTask(tasks *[]Task) {
	fmt.Print("Enter task description: ")
	reader := bufio.NewReader(os.Stdin)
	description, _ := reader.ReadString('\n')

	// Create a new task
	task := Task{
		Description: description[:len(description)-1],
		Completed:   false,
	}

	*tasks = append(*tasks, task)
	fmt.Println("Task added!")
}

func listTasks(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("No tasks found")
		return
	}

	for i, task := range tasks {
		status := "Not completed"
		if task.Completed {
			status = "Completed"
		}
		fmt.Printf("%d. %s [%s]\n", i+1, task.Description, status)
	}
}

func markTaskAsCompleted(tasks *[]Task) {
	fmt.Print("Enter task number to mark as completed: ")
	var index int
	fmt.Scanln(&index)

	if index <= 0 || index > len(*tasks) {
		fmt.Println("Invalid task number")
		return
	}
	// Lets mark the task as completed
	(*tasks)[index-1].Completed = true
	fmt.Println("Task marked as completed")
}

func removeTask(tasks *[]Task) {
	fmt.Print("Enter task number to remove: ")
	var index int
	fmt.Scanln(&index)

	if index <= 0 || index > len(*tasks) {
		fmt.Println("Invalid task number")
		return
	}

	// Remove the task; slice the array
	*tasks = append((*tasks)[:index-1], (*tasks)[index:]...)

	fmt.Println("Task Removed!")
}
