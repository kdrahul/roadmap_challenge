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
	"encoding/json"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
	"time"
)

const (
	Done       = "DONE"
	InProgress = "IN-PROGRESS"
	Todo       = "TODO"
	HELP       = `# Adding a new task
	task_tracker add "Buy groceries"
# Output: Task added successfully (ID: 1)

# Updating and deleting tasks
	task_tracker update 1 "Buy groceries and cook dinner"
	task_tracker delete 1

# Marking a task as in progress or done
	task_tracker mark-in-progress 1
	task_tracker mark-done 1

# Listing all tasks
	task_tracker list

# Listing tasks by status
	task_tracker list done
	task_tracker list todo
	task_tracker list in-progress`
)

// Colors
const (
	Reset = "\033[0m"
	Red   = "\033[31m"
)

type Task struct {
	Id        int       `json:"ID"`
	Title     string    `json:"Title"`
	State     string    `json:"State"`
	CreatedOn time.Time `json: "createdOn"`
	UpdatedOn time.Time `json: "updatedOn"`
}

func getAllTasks() {
	tasks, err := getDataFromFile()
	if err != nil {
		return
	}
	if len(tasks) <= 0 {
		fmt.Println(Red + "No data" + Reset)
		return
	}

	for _, task := range tasks {
		fmt.Printf("----------------------------------------\nID:\t%d\nTITLE:\t%s\nSTATUS:\t%s\nUPDATED ON:%v\n", task.Id, task.Title, task.State, task.UpdatedOn)
	}
}

func getMaxId(tasks []*Task) int {

	if len(tasks) != 0 {
		max := slices.MaxFunc(tasks, func(a, b *Task) int {
			if a.Id > b.Id {
				return 1
			} else if a.Id < b.Id {
				return -1
			}
			return 0
		})
		// fmt.Printf("Max Item: %v; ID: %v\n", max, max.Id)
		return max.Id
	}
	return 0
}

func addTask(t string) {
	tasks, err := getDataFromFile()
	if err != nil {
		fmt.Println("Couldn't get data:", err)
		return
	}
	newTask := Task{
		Id:        getMaxId(tasks) + 1,
		Title:     t,
		State:     Todo,
		CreatedOn: time.Now(),
		UpdatedOn: time.Now(),
	}
	fmt.Printf("Task added successfully: %d", newTask.Id)
	tasks = append(tasks, &newTask)
	writeToFile(tasks)
}

func getAllState(state string) {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		return
	}
	if string(data) == "" {
		fmt.Println(Red + "No data" + Reset)
		return
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Errorf("Error unmarshalling data from file: %v", err)
		return
	}
	var doneTasks []Task
	for _, task := range tasks {
		if strings.EqualFold(task.State, state) {
			doneTasks = append(doneTasks, task)
		}
	}
	fmt.Println(doneTasks)
}

func getDataFromFile() ([]*Task, error) {
	data, err := os.ReadFile(FILENAME)
	if err != nil {
		return nil, err
	}
	var tasks []*Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Errorf("Error unmarshalling data from file: %v", err)
		return nil, err
	}
	return tasks, nil
}

func writeToFile(tasks []*Task) {
	taskJson, err := json.Marshal(tasks)
	if err != nil {
		fmt.Println("Error marshalling data: %v", err)
		return
	}
	err = os.WriteFile(FILENAME, taskJson, 0644)
	if err != nil {
		fmt.Println(err)
	}
}

func markInProgress(id int) {

	tasks, err := getDataFromFile()
	if err != nil {
		fmt.Println("Couldn't get data:", err)
		return
	}

	for _, task := range tasks {
		if task.Id == id {
			task.State = InProgress
			task.UpdatedOn = time.Now()
		}
	}

	writeToFile(tasks)

}
func markDone(id int) {

	tasks, err := getDataFromFile()
	if err != nil {
		fmt.Println("Couldn't get data:", err)
		return
	}

	for _, task := range tasks {
		if task.Id == id {
			task.State = Done
			task.UpdatedOn = time.Now()
		}
	}

	writeToFile(tasks)

}

func deleteItem(i int) {
	tasks, err := getDataFromFile()
	if err != nil {
		fmt.Println("Couldn't get data:", err)
		return
	}

	var index int
	for idx, task := range tasks {
		if task.Id == i {
			index = idx
			break
		}
	}

	tasks = append(tasks[:index], tasks[index+1:]...)

	writeToFile(tasks)

}

func updateItem(i int, taskName string) {
	tasks, err := getDataFromFile()
	if err != nil {
		fmt.Println("Couldn't get data:", err)
		return
	}

	var index int
	for idx, task := range tasks {
		if task.Id == i {
			index = idx
			break
		}
	}

	tasks[index].Title = taskName
	tasks[index].UpdatedOn = time.Now()

	fmt.Printf("Task updated successfully: %d", tasks[index].Id)
	writeToFile(tasks)
}

const FILENAME = "store.json"

func main() {

	argsLength := len(os.Args)

	switch argsLength {
	case 2:
		switch os.Args[1] {
		case "list":
			getAllTasks()
		default:
			fmt.Println(HELP)
		}
	case 3:
		switch os.Args[1] {
		case "list":
			getAllState(os.Args[2]) // List tasks by status
		case "add":
			addTask(os.Args[2])
		case "delete":
			i, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Printf("Bad input, gimme a number: %v\n", err)
				return
			}
			deleteItem(i)
		case "mark-done":

			i, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Printf("Bad input, gimme a number: %v\n", err)
				return
			}
			markDone(i)
		case "mark-in-progress":
			i, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Printf("Bad input, gimme a number: %v\n", err)
				return
			}
			markInProgress(i)
		default:
			fmt.Println(HELP)
		}
	case 4:
		switch os.Args[1] {
		case "update":
			i, err := strconv.Atoi(os.Args[2])
			if err != nil {
				fmt.Printf("Bad input, gimme a number: %v\n", err)
				return
			}
			updateItem(i, os.Args[3])

		}

	default:
		fmt.Println(HELP)
	}

}
