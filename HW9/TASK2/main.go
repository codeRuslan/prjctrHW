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
}

var adminUser = User{
	Username: "admin",
	Password: "admin",
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
		GetStudents(w, r, students, adminUser)
	}))).
		Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", myRouter))
}

func GetStudents(w http.ResponseWriter, r *http.Request, students []Student, currentUser User) {
	vars := mux.Vars(r)
	key, _ := strconv.Atoi(vars["id"])

	for _, student := range students {
		if student.Id == int(key) {
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

		if username != adminUser.Username || password != adminUser.Password {
			w.WriteHeader(http.StatusUnauthorized)
			return
		}
		
		ctx := context.WithValue(r.Context(), "user", adminUser)
		
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

func setUserContext(r *http.Request, user User) *http.Request {
	ctx := context.WithValue(r.Context(), "user", user)
	return r.WithContext(ctx)
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
	return user.Username == adminUser.Username && user.Password == adminUser.Password
}

func isTeacherOfStudent(teacher User, student Student) bool {
	// For simplicity, let's assume the teacher's username should match the student's name.
	return teacher.Username == student.Name
}
