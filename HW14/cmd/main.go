package main

import (
	"HW15/internal/book"
	"HW15/internal/database"
	"HW15/internal/handler"
	"golang.org/x/exp/slog"
	"net/http"
	"os"
)

type App struct {
	Name    string
	Version string
}

func (app *App) Run() error {
	slog.Info("Starting up the API")

	var err error
	db, err := database.NewDatabase()
	if err != nil {
		return err
	}

	err = database.MigrateDB(db)
	if err != nil {
		slog.Error("Failed to set up database")
		return err
	}

	bookService := book.NewService(db)
	handler := handler.NewHandler(bookService)
	handler.SetupRoutes()

	if err := http.ListenAndServe(":8080", handler.Router); err != nil {
		slog.Error("Failed to set up the server")
		return err
	}

	return nil
}

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	app := App{
		Name:    "Book API",
		Version: "1.0",
	}

	if err := app.Run(); err != nil {
		slog.Error("Error starting up the App: ", err)
	}
}
