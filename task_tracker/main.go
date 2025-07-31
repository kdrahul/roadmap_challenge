// Task List used to track tasks
//
// Stored in JSON file
package main

// TODO: Add, Update, and Delete tasks
// TODO: Mark a task as in progress or done
// TODO: List all tasks
// TODO: List all tasks that are done
// TODO: List all tasks that are not done
// TODO: List all tasks that are in progress

// ---- Constraints ----
// - Use positional arguments in command line to accept user inputs.
// - Use a JSON file to store the tasks in the current directory.
// - The JSON file should be created if it does not exist.
// - Use the native file system module of your programming language to interact with the JSON file.
// - Do not use any external libraries or frameworks to build this project.
// - Ensure to handle errors and edge cases gracefully.

import (
	"fmt"
	"os"
	"encoding/json"
)

type State string

const (
	Done       State = "DONE"
	InProgress State = "IN-PROGRESS"
	Todo       State = "TODO"
)

// Colors
const (
	Reset = "\033[0m"
	Red = "\033[31m"
)

type Task struct {
	id int `json:"ID"`
	title string `json:"Title"`
	state State `json:"State"`
}


func getAllTasks() []Task{
			var data []byte
			f.Read(data)
			if string(data) == "" {
				fmt.Println(Red + "No data" + Reset)
				return
			}

			var tasks []Task
			err := json.Unmarshal(data, &tasks)
			if err != nil {
				fmt.Errorf("Error unmarshalling data from file: %v", err)
				return
			}
			fmt.Println(tasks)
			return tasks
}

func addTask(t string) {
	tasks := getAllTasks()
	newTask := Task{
		title: t,
		state: "todo"
	}
	append(tasks, newTask)
}

var f *os.File

func init() {
	filename := "store.json"
	file,err := os.Open(filename)
	if err != nil {
		fmt.Errorf("Error opening a file: %v\n", err)
		f, err = os.Create(filename)
		if err != nil {
			fmt.Errorf("Error in creating a file: %v\n", err)
		}
	}
	f = file
}

func main() {


	argsLength := len(os.Args)

	switch argsLength {
	case 2:
		switch os.Args[1] {
		case "list":
			getAllTasks(file)
		default:
			fmt.Println("You're using it wrong")
		}
	case 3:
		switch os.Args[1] {
		case "list":
			switch os.Args[2] {
			case "done":
			case "todo":
			case "in-progress":
			default:
				// TODO: Show help
			}
		case "add":
				addTask(os.Args[2])
		case "delete":
		case "mark-done":
			fmt.Println("Marking Done")
		case "mark-in-progress":
			fmt.Println("Good start")
		default:
			// TODO: Show help
		}
	case 4:
		switch os.Args[1] {
		case "update":
			fmt.Printf("Updating: %v -> %s", os.Args[2], os.Args[3])
		}

	default:
		fmt.Println("Here's how to use it")

	}

}
