package main

import (
	"fmt"

	"github.com/VolcharaWin/task-service/task"
)

func main() {
	fmt.Println("1")
	t := task.CreateTask()
	task.UpdateTask(&t, "test test", "2025-02-11T10:22:15Z")
}
