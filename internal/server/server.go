package server

import (
	"log"
	"net/http"
	"time"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/handlers"
)

type Server struct {
	Logger *log.Logger
	Server *http.Server
}

func MakeServer(logger *log.Logger) *Server {
	// Создаём роутер
	mux := http.NewServeMux()

	// Регистрируем хендлеры
	mux.HandleFunc("/upload", handlers.UploadHandle)
	mux.HandleFunc("/", handlers.MainHandle)

	// Настраиваем сервер
	s := &http.Server{
		Addr: ":8080",
		Handler: mux,
		ErrorLog: logger,
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout: 15 * time.Second,
	}

	server := Server{
		Logger: logger,
		Server: s,
	}

	return &server
}
