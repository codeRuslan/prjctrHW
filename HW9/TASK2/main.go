package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

type Student struct {
	Id          int    `json:"Id"`
	Name        string `json:"Name"`
	AverageMark int    `json:"Average Mark"`
}

type User struct {
	Username string
	Password string
}

var adminUser = User{
	Username: "admin",
	Password: "admin",
}

func main() {
	// Initialize students
	students := []Student{
		Student{1, "Ruslan", 99},
		Student{2, "Luiz", 89},
		Student{3, "John", 32},
		Student{4, "Alexa", 11},
	}

	// Create router
	myRouter := mux.NewRouter().StrictSlash(true)

	// Add route with authentication middleware
	myRouter.Handle("/student/{id}", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetStudents(w, r, students)
	}))).
		Methods("GET")

	// Start server
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func GetStudents(w http.ResponseWriter, r *http.Request, students []Student) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, student := range students {
		if student.Id == int(key) {
			json.NewEncoder(w).Encode(student)
			return
		}
	}

	// Student not found
	w.WriteHeader(http.StatusNotFound)
}

func auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		username, password, ok := r.BasicAuth()

		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		if username != adminUser.Username || password != adminUser.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		next.ServeHTTP(w, r)
	})
}
