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
		Task{"28.06.2023", "Do cleaning"},
		Task{"27.06.2023", "Do Homework"},
	}

	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/tasks", func(w http.ResponseWriter, r *http.Request) {
		GetAllTasks(w, r, tasks)
	})
	myRouter.HandleFunc("/task/{date}", func(w http.ResponseWriter, r *http.Request) {
		GetSingleTask(w, r, tasks)
	})

	log.Fatal(http.ListenAndServe(":10000", myRouter))
}

func GetAllTasks(w http.ResponseWriter, r *http.Request, tasks []Task) {
	json.NewEncoder(w).Encode(tasks)
}

func GetSingleTask(w http.ResponseWriter, r *http.Request, tasks []Task) {
	vars := mux.Vars(r)
	key := vars["date"]

	for _, task := range tasks {
		if task.Date == key {
			json.NewEncoder(w).Encode(task)
			return
		}
	}

	// Task not found
	w.WriteHeader(http.StatusNotFound)
}
