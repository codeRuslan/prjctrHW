package handler

import (
	"HW15/config"
	"HW15/internal/book"
	"encoding/json"
	"errors"
	"github.com/gorilla/mux"
	"golang.org/x/exp/slog"
	"net"
	"net/http"
	"strings"
	"time"
)

type Response struct {
	Message string
	Error   string
}

type Handler struct {
	Router  *mux.Router
	Service *book.Service
}

func NewHandler(service *book.Service) *Handler {
	return &Handler{
		Service: service,
	}
}

func (h *Handler) SetupRoutes() {
	slog.Info("Setting up Routes")
	h.Router = mux.NewRouter()
	h.Router.Use(logMW)
	h.Router.HandleFunc("/api/book/{id}", h.GetBook).Methods("GET")
	h.Router.HandleFunc("/api/book", h.GetAllBooks).Methods("GET")
	h.Router.HandleFunc("/api/book/{id}", h.PostBook).Methods("PUT")
	h.Router.HandleFunc("/api/book/{id}", h.DeleteBook).Methods("DELETE")
	h.Router.HandleFunc("/api/health", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(http.StatusOK)
		if err := json.NewEncoder(w).Encode(Response{Message: "I am Alive!"}); err != nil {
			panic(err) // тут паніка, бо якщо цей endpoint не працює - отже нічого не працює
		}
	})
}

func logMW(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		slog.Info(r.RequestURI)

		startedAt := time.Now()
		defer func() {
			elapsed := time.Since(startedAt)
			slog.Info(r.RequestURI, "elapsed", elapsed)
		}()

		next.ServeHTTP(w, r)
	})
}

func ApiKeyMiddleware() func(handler http.Handler) http.Handler {
	apiKeyHeader := conf.APIKeyHeader
	apiKeys := conf.APIKeys

	reverseKeyIndex := make(map[string]string)
	for name, key := apiKeys {
		reverseKeyIndex[key] = name
	}

	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			apiKey, err := bearerToken(r, apiKeyHeader)

			if err != nil {
				sendErrorResponse(w, "request failed API key authentication", err)
				return
			}

			_, found := reverseKeyIndex[apiKey]

			if !found {
				_, _, err := net.SplitHostPort(r.RemoteAddr)
				if err != nil {
					sendErrorResponse(w, "Failed to parse remote address", err)
					return
				}
			}

			next.ServeHTTP(w ,r)
		})
	}
}

func bearerToken (r *http.Request, header string) (string, error) {
	rawToken := r.Header.Get(header)
	pieces := strings.SplitN(rawToken, " ", 2)

	if len(pieces) < 2 {
		return "", errors.New("Token with incorrect bearer format")
	}

	token := strings.TrimSpace(pieces[1])

	return token, nil
}
