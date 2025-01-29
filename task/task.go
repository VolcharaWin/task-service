package task

import (
	"encoding/json"
	"fmt"
	"os"
	"time"
)

type Task struct {
	Content string    `json:"content"`
	Active  bool      `json:"active"`
	TimeIn  time.Time `json:"timein"`
	TimeOut time.Time `json:"timeout"`
}

func CreateTask() Task {
	jsonTask, err := os.ReadFile("../task/task.json")
	if err != nil {
		panic("could not open task.json")
	}

	var t Task

	err = json.Unmarshal(jsonTask, &t)

	if err != nil {
		panic("couldn't unmarshal the jsonTask")
	}
	return t
}

func DeleteTask(t *Task) {
	t.Active = false
}

// Update task function which recieves changes of content or/and date of completion
func UpdateTask(t *Task, args ...string) {
	if len(args) > 0 {
		fmt.Println("зашел в contentFlag")
		t.Content = args[0]
	}
	if len(args) > 1 {
		fmt.Println("зашел в dateFlag")
		newDateOut, err := time.Parse(time.RFC3339, args[1])
		fmt.Println("newDateOut: ", newDateOut)
		if err != nil {
			fmt.Println("err updateFlag with dateFlag: ", err)
			return
		}
		t.TimeOut = newDateOut
	}
}
