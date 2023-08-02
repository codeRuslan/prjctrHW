package handler

import (
	"HW15/internal/book"
	"encoding/json"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net/http"
	"strconv"
)

func sendErrorResponse(w http.ResponseWriter, message string, err error) {
	w.WriteHeader(http.StatusInternalServerError)
	if err := json.NewEncoder(w).Encode(Response{Message: message, Error: err.Error()}); err != nil {
		panic(err)
	}
}

func (h *Handler) GetAllBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	slog.With("operation", "handler.book_methods.GetAllBooks").Info("Starting operation")
	defer slog.With("operation", "handler.book_methods.GetBook").Info("Finished operation")

	books, err := h.Service.GetAllComments()

	if err != nil {
		sendErrorResponse(w, "Failed to retrieve all the comments", err)
	}

	if err := json.NewEncoder(w).Encode(books); err != nil {
		sendErrorResponse(w, "Failed to encode message", err)
	}
}

func (h *Handler) GetBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	slog.With("operation", "handler.book_methods.GetBook").Info("Starting operation")
	defer slog.With("operation", "handler.book_methods.GetBook").Info("Finished operation")

	vars := mux.Vars(r)

	id := vars["id"]
	i, err := strconv.ParseUint(id, 10, 64)
	if err != nil {
		sendErrorResponse(w, "Unable to parse UINT from ID", err)
	}

	book, err := h.Service.GetComment(uint(i))
	if err != nil {
		sendErrorResponse(w, "Error retrieving Comment by ID", err)
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		sendErrorResponse(w, "Failed to encode message", err)
	}
}

func (h *Handler) PostBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	slog.With("operation", "handler.book_methods.PostBook").Info("Starting operation")
	defer slog.With("operation", "handler.book_methods.PostBook").Info("Finished operation")

	var book book.Book

	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		sendErrorResponse(w, "Failed to decode JSON body", err)
	}

	book, err := h.Service.PostComment(book)
	if err != nil {
		sendErrorResponse(w, "Failed to post new book", err)
	}

	if err := json.NewEncoder(w).Encode(book); err != nil {
		sendErrorResponse(w, "Failed to encode message", err)
	}
}

func (h *Handler) DeleteBook(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	slog.With("operation", "handler.book_methods.DeletetBook").Info("Starting operation")
	defer slog.With("operation", "handler.book_methods.DeleteBook").Info("Finished operation")

	vars := mux.Vars(r)
	id := vars["id"]
	bookID, err := strconv.ParseUint(id, 10, 64)

	if err != nil {
		sendErrorResponse(w, "Failed to parse Uint from ID", err)
	}

	err = h.Service.DeleteBook(uint(bookID))
	if err != nil {
		sendErrorResponse(w, "Failed to delete comment by comment iD", err)
	}

	if err := json.NewEncoder(w).Encode(Response{Message: "Successfully Deleted"}); err != nil {
		sendErrorResponse(w, "Failed to encode message", err)
	}

}
