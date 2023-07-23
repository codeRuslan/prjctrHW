package router

import (
	"awesomeProject3/internal/dependencies"
	"awesomeProject3/internal/handler"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

func HandleRequests() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/users", handler.AllUsers).Methods("GET")
	myRouter.HandleFunc("/creteUser", handler.CreateUser).Methods("POST")
	myRouter.HandleFunc("/deleteUser", handler.DeleteUser).Methods("POST")
	log.Fatal(http.ListenAndServe(":"+dependencies.Cfg.Server.Port, myRouter))
}
