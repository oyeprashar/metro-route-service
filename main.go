package main

import (
	"github.com/go-chi/chi/v5"
	"net/http"
	"time"
)

/*
	TODO:
		1. Set up monitoring and altering via prometheus and grafana
*/

func main() {

	// using chi to route the incoming calls to our server
	r := chi.NewRouter()

	// TODO:

	// creating an HTTP server and listening to 8080, chi handles the incoming requests
	server := &http.Server{
		Addr:              ":8080",
		ReadHeaderTimeout: 3 * time.Minute,
		Handler:           r,
	}

	err := server.ListenAndServe()

	if err != nil {
		panic(err)
	}
}
