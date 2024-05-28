package main

import (
	"fmt"
)

type Task struct {
	id  int
	title string
	description string
	status string


}

func newTask(id int, title string, description string, status string) Task {
	return Task{
		id: id,
		title: title,
		description: description,
		status: status,
	
	}
}

type TaskManager struct {
	tasks []Task
}

func newTaskManager() *TaskManager {
	return &TaskManager{
		tasks: make([]Task, 0),
	}
}

func (tm *TaskManager) addTask(task Task) {
	tm.tasks = append(tm.tasks, task)
}

func main() {
	taskManager := newTaskManager()

	task := newTask(1, "Learn Go", "Michael is Learning go", "pending")

	taskManager.addTask(task)
	
	fmt.Printf("Task Manager: %v\n", *taskManager)
}