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
	"time"
)

type Status struct {
	NotDone    bool
	InProgress bool
	Done       bool
}

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
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
		Status: Status{
			NotDone:    true,
			InProgress: false,
			Done:       true,
		},
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}

	return append(task, tasks)
}

func loadJSON() []Task {
	data, err := os.ReadFile("task.json")
	if err != nil {
		_ = fmt.Errorf("error to read the file")
	}

	var tasks []Task
	err = json.Unmarshal(data, &tasks)
	if err != nil {
		_ = fmt.Errorf("error to read the file")
	}
	return tasks
}

func saveJSON() []Task {
	var tasks []Task
	data, err := json.Marshal(tasks)
	if err != nil {
		_ = fmt.Errorf("error to read the file")
	}

	err = os.WriteFile("task.json", data, 0644)
	if err != nil {
		_ = fmt.Errorf("error to read the file")
	}

	return tasks
}

func main() {
	loadJSON()
	saveJSON()
	log.Print("Use task-cli + command to use this app...")
	if len(os.Args) < 3 {
		fmt.Println("\nInvalid arguments!!!" +
			"\nUse command like task-cli add or task-cli show.")
		return
	}

	switch os.Args[1] {
	case "add":
		//AddTask(,os.Args[3])
	}
}
