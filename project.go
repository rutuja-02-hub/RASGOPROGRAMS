package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type TaskList struct {
	Tasks []Task `json:"tasks"`
}

const fileName = "tasks.json"

func loadTasks() TaskList {
	var taskList TaskList
	file, err := ioutil.ReadFile(fileName)
	if err == nil {
		json.Unmarshal(file, &taskList)
	}
	return taskList
}

func saveTasks(taskList TaskList) {
	file, _ := json.MarshalIndent(taskList, "", "  ")
	ioutil.WriteFile(fileName, file, 0644)
}

func addTask(name string) {
	taskList := loadTasks()
	newTask := Task{ID: len(taskList.Tasks) + 1, Name: name}
	taskList.Tasks = append(taskList.Tasks, newTask)
	saveTasks(taskList)
	fmt.Println("Task added!")
}

func listTasks() {
	taskList := loadTasks()
	if len(taskList.Tasks) == 0 {
		fmt.Println("No tasks available.")
		return
	}
	for _, task := range taskList.Tasks {
		fmt.Printf("%d. %s\n", task.ID, task.Name)
	}
}

func deleteTask(id int) {
	taskList := loadTasks()
	newTasks := []Task{}
	for _, task := range taskList.Tasks {
		if task.ID != id {
			newTasks = append(newTasks, task)
		}
	}
	taskList.Tasks = newTasks
	saveTasks(taskList)
	fmt.Println("Task deleted!")
}

func main() {
	var choice int
	for {
		fmt.Println("1. Add Task")
		fmt.Println("2. List Tasks")
		fmt.Println("3. Delete Task")
		fmt.Println("4. Exit")
		fmt.Print("Enter choice: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			var taskName string
			fmt.Print("Enter task: ")
			fmt.Scan(&taskName)
			addTask(taskName)
		case 2:
			listTasks()
		case 3:
			var taskID int
			fmt.Print("Enter task ID to delete: ")
			fmt.Scan(&taskID)
			deleteTask(taskID)
		case 4:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid choice! Try again.")
		}
	}
}
