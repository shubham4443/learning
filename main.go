package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
)

// Server is the object which sets up the server and handles all server operations.
type Server struct {
	store  map[string]interface{}
	router *mux.Router
	mu     sync.Mutex
}

func main() {
	s := &Server{
		store:  make(map[string]interface{}),
		router: mux.NewRouter(),
	}
	s.routes()
	fmt.Println("Server listening....")
	log.Fatal(http.ListenAndServe(":9000", s.router))
}
