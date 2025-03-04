package main

import (
	"flag"
	"fmt"
)

type Task struct {
	Id          int    `json:"id"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
}

var (
	command     string
	description string
	id          int
)

func main() {
	flag.StringVar(&command, "command", "", "Команда для выполнения (add, list, done, remove, clear)")
	flag.StringVar(&description, "description", "", "Описание задачи")
	flag.IntVar(&id, "id", 0, "Идентификатор задачи")

	flag.Parse()

	switch command {
	case "add":
		fmt.Println("Add")
	case "list":
		fmt.Println("List")
	case "done":
		fmt.Println("Done")
	case "remove":
		fmt.Println("Remove")
	case "clear":
		fmt.Println("Clear")
	default:
		fmt.Println("Нужно указать команду из этого списка (add, list, done, remove, clear)")
	}
}

func loadTasks() {

}

func saveTasks() {

}

func addTask() {

}

func listTasks() {

}

func completeTask() {

}

func removeTask() {

}

func clearCompletedTasks() {

}
