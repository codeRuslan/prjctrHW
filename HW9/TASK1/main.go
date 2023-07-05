package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

type Task struct {
	Date     string `json:"Date"`
	TaskName string `json:"Task Name"`
}

func main() {
	tasks := []Task{
		Task{"30.06.2023", "Wash dishes"},
		Task{"29.06.2023", "Do homework"},
		Task{"29.06.2023", "Do cleaning"},
		Task{"27.06.2023", "Do Homework"},
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		GetAllTasks(w, r, tasks)
	}).Methods("GET")
	myRouter.HandleFunc("/task/{date}", func(w http.ResponseWriter, r *http.Request) {
		GetSingleTask(w, r, tasks)
	}).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetAllTasks(w http.ResponseWriter, r *http.Request, tasks []Task) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(tasks)
}

func GetSingleTask(w http.ResponseWriter, r *http.Request, tasks []Task) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["date"]
	var foundTasks []Task

	for _, task := range tasks {
		if task.Date == key {
			foundTasks = append(foundTasks, task)
		}
	}

	if len(foundTasks) == 0 {
		// Task not found
		w.WriteHeader(http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(foundTasks)
}
