package main

import (
	"log"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/internal/server"
)

func main() {
	logger := log.Logger{}
	s := server.MakeServer(&logger)

	log.Fatal(s.Server.ListenAndServe())
}
