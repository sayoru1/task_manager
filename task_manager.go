package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"strconv"
	"strings"
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

func saveTask(tasks []Task, file *os.File) error {
	_, err := file.Seek(0, 0)
	if err != nil {
		fmt.Println(err)
		return err
	}
	err = file.Truncate(0)
	if err != nil {
		fmt.Println(err)
		return err
	}

	writer := bufio.NewWriter(file)

	for _, t := range tasks {
		_, err := fmt.Fprintf(writer, "%d|%s|%s|%t\n", t.ID, t.Title, t.Description, t.Completed)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}
	return writer.Flush()
}

func loadTask(file *os.File) ([]Task, error) {
	reader := bufio.NewReader(file)

	var tasks []Task

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err == io.EOF {
				if line == "" {
					break
				}
			} else {
				fmt.Println(err)
				return tasks, err
			}
		}
		line = strings.TrimSpace(line)
		if line == "" {
			if err == io.EOF {
				break
			}
			continue
		}

		parts := strings.Split(line, "|")

		id, err := strconv.Atoi(parts[0])
		if err != nil {
			fmt.Println(err)
			return tasks, err
		}

		completed, err := strconv.ParseBool(parts[3])
		if err != nil {
			fmt.Println(err)
			return tasks, err
		}

		t := Task{id, parts[1], parts[2], completed}
		tasks = append(tasks, t)
		if err == io.EOF {
			break
		}
	}
	return tasks, nil
}

func main() {

	home, err := os.UserHomeDir() //iniatilization of home directory variable
	if err != nil {
		fmt.Println(err)
		return
	}
	path := filepath.Join(home, "Documents", "tasks.txt") //iniatilization of path variable

	file, err := os.OpenFile(path, os.O_RDWR|os.O_CREATE, 0644) //file creation

	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	tasks, err = loadTask(file)

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
			err = saveTask(tasks, file)
			if err != nil {
				fmt.Println(err)
			}
		case 3:
			_, err = completeTask(tasks)
			if err != nil {
				fmt.Println(err)
				return
			}

			err = saveTask(tasks, file)
			if err != nil {
				fmt.Println(err)
				return
			}
		case 4:
			tasks, err = deleteTask(tasks)
			if err != nil {
				fmt.Println(err)
				return
			}
			err = saveTask(tasks, file)
			if err != nil {
				fmt.Println(err)
			}
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
