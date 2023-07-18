package main

import (
	"context"
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
	Role     string
}

var users = []User{
	User{Username: "admin", Password: "admin", Role: "admin"},
	User{Username: "teacher1", Password: "teacher1", Role: "teacher"},
	User{Username: "teacher2", Password: "teacher2", Role: "teacher"},
}

func main() {
	students := []Student{
		Student{1, "Ruslan", 99},
		Student{2, "Luiz", 89},
		Student{3, "John", 32},
		Student{4, "Alexa", 11},
	}

	myRouter := mux.NewRouter().StrictSlash(true)

	myRouter.Handle("/student/{id}", auth(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		GetStudents(w, r, students)
	}))).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func GetStudents(w http.ResponseWriter, r *http.Request, students []Student) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, student := range students {
		if student.Id == int(key) {
			currentUser := getCurrentUser(r)
			if isAdmin(currentUser) || isTeacherOfStudent(currentUser, student) {
				json.NewEncoder(w).Encode(student)
				return
			} else {
				w.WriteHeader(http.StatusUnauthorized)
				return
			}
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

		var currentUser User
		for _, user := range users {
			if user.Username == username && user.Password == password {
				currentUser = user
				break
			}
		}

		if currentUser.Role == "" {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user", currentUser)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func getCurrentUser(r *http.Request) User {
	user := r.Context().Value("user")
	if user != nil {
		if currentUser, ok := user.(User); ok {
			return currentUser
		}
	}
	return User{}
}

func isAdmin(user User) bool {
	return user.Role == "admin"
}

func isTeacherOfStudent(teacher User, student Student) bool {
	// For simplicity, let's assume the teacher's username should match the student's name.
	return teacher.Role == "teacher" && teacher.Username == student.Name
}
