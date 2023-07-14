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
	myRouter.HandleFunc("/tasks", GetAllTasks(tasks)).Methods("GET")
	myRouter.HandleFunc("/task/{date}", GetSingleTask(tasks)).Methods("GET")

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetAllTasks(tasks []Task) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(tasks)
	}
}

func GetSingleTask(tasks []Task) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
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
}
