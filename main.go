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
	"os"
	"path/filepath"
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
		Status:      "NotDone",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	return append(task, tasks)
}

func showTask() {

	/*	if len(tsk) == 0 {
			fmt.Println("No tasks found.")
			return
		}
		for _, t := range tsk {
			fmt.Printf("[%d] %s | %s\n", t.ID, t.Description, t.Status)
		}
		i can use for show function
	*/
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

	if len(os.Args) <= 1 {
		fmt.Println("Use task-cli + command to use this app...")
	}
	//verify if the arguments to task are valid
	if len(os.Args) < 2 {
		fmt.Println("\nInvalid arguments!!!" +
			"\nUse command like task-cli add or task-cli show.")
		return
	}

	switch os.Args[1] {
	case "add":
		if len(os.Args) < 3 {
			fmt.Println("Usage: task-cli add \"task description\"")
			return
		}
		tsk = AddTask(tsk, os.Args[2])
		saveJSON(tsk)
		fmt.Printf("Task added (ID %d): %s\n", tsk[len(tsk)-1].ID, os.Args[2])

	default:
		fmt.Printf("Unknown command: %s\n", os.Args[1])
	}
}
