/*
- create struct for task
- func for interact with the task
	- show all
	- add, edit,create, delete
	- show filtered task(done, in progress, not done
- store the created task in JSON file

-----------------------------------
json.Unmarshal(data, &tasks)JSON -> struct
json.Marshal(tasks) struct -> JSON
*/

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"createdAt"`
	UpdatedAt   time.Time `json:"updatedAt"`
}

func AddTask(task []Task, description string) []Task {
	newId := 1

	if len(task) > 0 {
		bigrID := task[0].ID
		for _, t := range task {
			if t.ID > bigrID {
				bigrID = t.ID
			}
		}
		newId = bigrID + 1
	}

	tasks := Task{
		ID:          newId,
		Description: description,
		Status:      "not done",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return append(task, tasks)
}

func DeleteTask(task []Task, idToDel int) ([]Task, bool) {
	res := []Task{}
	found := false
	for _, t := range task {
		if t.ID == idToDel {
			found = true
		} else {
			res = append(res, t)
		}
	}
	return res, found
}

func ShowTask(task []Task) {

	if len(task) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for _, t := range task {
		fmt.Printf("[%d] %s | %s\n", t.ID, t.Description, t.Status)
	}
}

func ChangeStatus(task []Task, id int, newStatus string) {
	if len(task) == 0 {
		fmt.Println("No tasks found.")
		return
	}
	for i := range task {
		if task[i].ID == id {
			task[i].Status = strings.ToLower(newStatus)
			return
		}
	}
	fmt.Println("Taks not found.")
}

func dataPath() string {
	exe, err := os.Executable()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error getting executable path: %v\n", err)
		os.Exit(1)
	}
	return filepath.Join(filepath.Dir(exe), "task.json")
}

func loadJSON() []Task {
	path := dataPath()
	data, err := os.ReadFile(path)
	if err != nil {
		if os.IsNotExist(err) {
			return []Task{}
		}
		fmt.Fprintf(os.Stderr, "Error reading %s: %v\n", path, err)
		return []Task{}
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error parsing %s: %v\n", path, err)
		return []Task{}
	}
	return tasks
}

func saveJSON(tasks []Task) {
	path := dataPath()
	data, err := json.Marshal(tasks)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error encoding tasks: %v\n", err)
		return
	}

	os.MkdirAll(filepath.Dir(path), 0755)
	err = os.WriteFile(path, data, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error writing %s: %v\n", path, err)
	}
}

func main() {
	tsk := loadJSON()
	saveJSON(tsk)

	//verify if the arguments to task are valid
	if len(os.Args) < 2 {
		fmt.Println("\nInvalid arguments!!!" +
			"\nUse command like task-cli add or task-cli show." +
			"\nOr use task-cli help to see the commands you can use")
		return
	}

	switch os.Args[1] {
	//add task
	case "add":
		if len(os.Args) < 2 {
			fmt.Println("Usage: task-cli add \"task description\"")
			return
		}
		tsk = AddTask(tsk, os.Args[2])
		saveJSON(tsk)
		fmt.Printf("Task added (ID %d): %s\n", tsk[len(tsk)-1].ID, os.Args[2])
	//delete task
	case "delete":
		if len(os.Args) < 2 {
			fmt.Println("Usage: task-cli add \"task description\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID. Must be a number.")
			return
		}
		_, found := DeleteTask(tsk, id)
		if !found {
			fmt.Println("The Taks doesn't exist, try delete an valid task.")
		}
		log.Printf("Task to delete (ID %d): %s\n", tsk[len(tsk)-1].ID, os.Args[2])
		tsk, _ = DeleteTask(tsk, id)
		saveJSON(tsk)
	//change the status
	case "change-status":
		//utilizar  argStatus = os.Args[3] task-cli change-status 1 "Done"
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add \"task description\"")
			return
		}
		id, err := strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Println("Invalid ID")
			return
		}
		ChangeStatus(tsk, id, os.Args[3])
		saveJSON(tsk)
	case "help":
		fmt.Printf("For add new tasks use: task-cli add \"Task name\" \n" +
			"For delete tasks use: task-cli delete \"Task id\" \n" +
			"For change the status of tasks use: task-cli change-status \"Task id\" \"The new status\" \n")
	case "update":

	case "list":

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
