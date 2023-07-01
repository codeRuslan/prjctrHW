package main

import (
	"encoding/json"
	"net/http"
)

var Tasks []Task

type Task struct {
	Date     string `json:"Date"`
	TaskName string `json:"Task Name"`
}

func main() {
	Tasks = []Task{
		Task{"30.06.2023", "Wash dishes"},
		Task{"29.06.2023", "Do homework"},
		Task{"28.06.2023", "Do cleaning"},
		Task{"27.06.2023", "Do Homework"},
	}
	handleRequests()
}

func handleRequests() {
	http.HandleFunc("/tasks", GetTasks)
	http.ListenAndServe(":8080", nil)
}

func GetTasks(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(Tasks)

}
