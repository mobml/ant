package tasks

import (
	"encoding/json"
	"fmt"
	"github.com/mobml/ant/utils"
	"os"
)

type Task struct {
	ID          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var tasks []Task
var TaskID int

func LoadTasks() error {

	path, err := utils.GetPath("tasks.json")

	if err != nil {
		return err
	}

	file, err := os.Open(path)

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
	path, err := utils.GetPath("tasks.json")

	if err != nil {
		return err
	}
	file, err := os.Create(path)

	if err != nil {
		return err
	}

	defer file.Close()

	encoder := json.NewEncoder(file)
	err = encoder.Encode(tasks)
	return err
}

func AddTask(description string) {
	TaskID++
	newTask := Task{
		ID:          TaskID,
		Description: description,
		Completed:   false,
	}

	tasks = append(tasks, newTask)
	fmt.Printf("Task '%s' has been added", newTask.Description)
	saveTasks()
}

func ListTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks to show")
	}

	for _, task := range tasks {
		fmt.Printf("%d. %s, [Completed: %v]\n",
			task.ID, task.Description, task.Completed)
	}
}

func DeleteTask(id int) {
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

func CompleteTask(id int) {
	for i, task := range tasks {
		if task.ID == id {
			tasks[i].Completed = true
			saveTasks()
			fmt.Printf("Task '%s' marked as completed\n", task.Description)
			return
		}
	}
	fmt.Printf("Task not founded\n")
}
