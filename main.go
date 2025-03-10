package main

import "fmt"

func main() {
	var input string
	loadTasks()
	fmt.Println("Ingresa tu tarea")
	fmt.Scan(&input)
	addTask(input)
	listTasks()
	deleteTask(1)
	listTasks()
	completeTask(3)
	listTasks()
}
