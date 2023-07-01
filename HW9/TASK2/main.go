package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

var Students []Student

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
	Students = []Student{
		Student{1, "Ruslan", 99},
		Student{2, "Luiz", 89},
		Student{3, "John", 32},
		Student{4, "Alexa", 11},
	}
	handleRequests()
}

func handleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.Handle("/student/{id}", auth(http.HandlerFunc(GetStudents))).Methods("GET")
	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, student := range Students {
		if student.Id == int(key) {
			json.NewEncoder(w).Encode(student)
		}
	}
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
