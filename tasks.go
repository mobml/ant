package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks []Task
var TaskID int

func loadTasks() error {
	file, err := os.Open("tasks.json")

	if err != nil {
		if os.IsNotExist(err) {
			return nil
		}
		return err
	}

	defer file.Close()
	decoder := json.NewDecoder(file)

	err = decoder.Decode(&tasks)
	if len(tasks) > 0 {
		TaskID = tasks[len(tasks)-1].ID
	}
	return err
}

func saveTasks() error {
	file, err := os.Create("tasks.json")

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	return err
}

func addTask(description string) {
	TaskID++
	newTask := Task{
		ID:          TaskID,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
	saveTasks()
}

func listTasks() {
	if len(tasks) == 0 {
		fmt.Println("No hay tareas que mostrar")
	}

	for _, task := range tasks {
		fmt.Printf("%d. %s, [Completed: %v]\n",
			task.ID, task.Description, task.Completed)
	}
}

func deleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks = append(tasks[:i], tasks[i+1:]...)
			saveTasks()
			fmt.Printf("Task '%s' removed\n", task.Description)
			return
		}
	}
	fmt.Printf("Task not founded\n")
}

func completeTask(id int) {
	for i, task := range tasks {
		if task.ID == id  {
			tasks[i].Completed = true
			saveTasks()
			fmt.Printf("Task '%s' marked as completed\n", task.Description)
			return
		}
	}
	fmt.Printf("Task not founded\n")
}
