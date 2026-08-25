package main

import (
	"errors"
	"fmt"
)

type Task struct {
	ID          int
	Title       string
	Description string
	Completed   bool
}

var tasks []Task

func addTask(tasks []Task, title string, description string) []Task {
	var t Task = Task{len(tasks) + 1, title, description, false}
	tasks = append(tasks, t)
	return tasks
}

func showTask(tasks []Task) {
	if len(tasks) == 0 {
		fmt.Println("Task list is empty")
		return
	}
	for index, i := range tasks {
		status := "[ ]"
		if i.Completed {
			status = "[X]"
		}
		fmt.Printf("ID: %d | %s %s - %s\n", index+1, status, i.Title, i.Description)
	}
}

func completeTask(tasks []Task) (Task, error) {
	fmt.Println("Choose your task")
	showTask(tasks)

	var choose int
	fmt.Scan(&choose)

	if choose > len(tasks) || choose <= 0 {
		return Task{}, errors.New("Task doen't exist")
	}

	tasks[choose-1].Completed = true
	return tasks[choose-1], nil
}

func deleteTask(tasks []Task) ([]Task, error) {
	fmt.Println("Choose your task")
	showTask(tasks)

	var choose int
	fmt.Scan(&choose)

	if choose > len(tasks) || choose <= 0 {
		return tasks, errors.New("Task doen't exist")
	}

	tasks = append(tasks[:choose-1], tasks[choose:]...)
	return tasks, nil
}

func main() {

	for {

		fmt.Println(`--- TASK MANAGER ---
1. Show Task
2. Add Task
3. Complete Task
4. Delete Task
5. Exit

Choose action:`)

		var action int
		var chsToCon string

		fmt.Scan(&action)
		switch action {
		case 1:
			showTask(tasks)
		case 2:
			fmt.Println("Write a title and description of your task")

			var title string
			var description string

			fmt.Println("Title:")
			fmt.Scan(&title)
			fmt.Println("Description:")
			fmt.Scan(&description)
			tasks = addTask(tasks, title, description)
		case 3:
			completeTask(tasks)
		case 4:
			tasks, _ = deleteTask(tasks)
		case 5:
			return
		default:
			fmt.Println("This action doesn't exist")

		}
		fmt.Println("Do you want to continue:Y/n")
		fmt.Scan(&chsToCon)
		if chsToCon != "y" {
			break
		}
	}
}
