package server

import (
	"log"
	"net/http"
	"time"

	"github.com/AidanFogarty/go-kubernetes-starter/cmd/handlers"
)

// StartServer starts the server
func StartServer() {

	srv := &http.Server{
		Handler:      handlers.HandleRequests(),
		Addr:         ":8080",
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	log.Printf("Server Starting At %s", srv.Addr)
	srv.ListenAndServe()

}
