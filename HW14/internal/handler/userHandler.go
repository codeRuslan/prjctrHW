package handler

import (
	"awesomeProject3/internal/dependencies"
	"awesomeProject3/internal/types"
	"encoding/json"
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/mattn/go-sqlite3"
	"net/http"
)

func AllUsers(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", dependencies.Cfg.Database.DatabaseLocation)
	if err != nil {
		fmt.Println(err.Error())
		panic("failed to connect to database")
	}

	defer db.Close()

	var users []types.User

	db.Find(&users)
	fmt.Println("{}", users)

	json.NewEncoder(w).Encode(users)
}

func CreateUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "/Users/ruslanpilipyuk/GolandProjects/awesomeProject3/cmd/test.db")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
	}

	defer db.Close()
	var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
	}
	db.Create(&user)

	fmt.Fprintf(w, "User Has Been Updated/Created Successfully")
}

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	db, err := gorm.Open("sqlite3", "/Users/ruslanpilipyuk/GolandProjects/awesomeProject3/cmd/test.db")
	if err != nil {
		fmt.Println(err.Error())
		http.Error(w, "Failed to connect to DB", http.StatusInternalServerError)
		return
	}

	defer db.Close()

	var user types.User
	err = json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		http.Error(w, "Failed to parse request body", http.StatusBadRequest)
		return
	}

	userName := user.Name
	userSecondName := user.SecondName

	if db.Where("name = ? and second_name = ?", userName, userSecondName).First(&user).RecordNotFound() {
		http.Error(w, "User Not Found", http.StatusNotFound)
		return
	}

	if err := db.Delete(&user).Error; err != nil {
		http.Error(w, "Failed to delte user", http.StatusInternalServerError)
		return
	}

	fmt.Fprintf(w, "User Has Been Successfully Deleted")
}
