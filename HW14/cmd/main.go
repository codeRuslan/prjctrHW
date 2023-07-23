package main

import (
	"awesomeProject3/internal/dependencies"
	"awesomeProject3/internal/router"
)

func main() {
	dependencies.InitialMigration()
	router.HandleRequests()
}
