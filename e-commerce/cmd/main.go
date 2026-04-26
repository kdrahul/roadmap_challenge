package main

import (
	"database/sql"
	"e-commerce/api"
	"e-commerce/db"
	"e-commerce/handler"
	"fmt"
	"log"
	"log/slog"
	"net"
	"os"
	"time"

	"net/http"

	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/chi/v5"
)

func InitLogs() {
	slog.SetDefault(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{})))
}

var DB *sql.DB

const (
	HOST = "0.0.0.0"
	PORT = "9123"
)

func DefaultRequestErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Print(err)
	http.Error(w, fmt.Sprintf(`{"message": "bad request", "error": "%s", "statusCode": "400"}`, err.Error()), 400)
}
func DefaultResponseErrorHandler(w http.ResponseWriter, r *http.Request, err error) {
	log.Println(err)
	http.Error(w, fmt.Sprintf(`{"message": "internal server error", "error": "%s", "statusCode": "500"}`, err.Error()), http.StatusInternalServerError)
}

func main() {

	sqlite := db.OpenSqlite()

	server := handler.NewServer(&sqlite)

	r := chi.NewMux()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Use(middleware.Timeout(30 * time.Second))

	strictHandler := api.NewStrictHandlerWithOptions(server, nil, api.StrictHTTPServerOptions{
		RequestErrorHandlerFunc:  DefaultRequestErrorHandler,
		ResponseErrorHandlerFunc: DefaultResponseErrorHandler,
	})
	r.Route("/api", func(r chi.Router) {
		api.HandlerFromMux(strictHandler, r)
	})

	s := http.Server{
		Addr:         net.JoinHostPort("0.0.0.0", "9123"),
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 30 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	if err := s.ListenAndServe(); err != nil {
		log.Fatal(err)
	}
	log.Print("Server stopped")
}
