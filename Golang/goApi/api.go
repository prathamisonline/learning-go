package main

import (
	"fmt"
	"net/http"
	"strings"
)

type APIServer struct {
	addr string
}

// Corrected function name to NewAPIServer
func NewAPIServer(addr string) *APIServer {
	return &APIServer{addr: addr} // Returns a pointer to a new APIServer instance
}

func (s *APIServer) Run() error {
	router := http.NewServeMux()
	router.HandleFunc("/users/", func(w http.ResponseWriter, r *http.Request) {
		// Extract userID from the URL
		userId := strings.TrimPrefix(r.URL.Path, "/users/")
		w.Write([]byte("User ID: " + userId)) // Corrected response writing
	})

	server := http.Server{
		Addr:    s.addr,
		Handler: router,
	}
	fmt.Printf("Starting server on %s\n", s.addr)
	// This is the line that starts the server
	return server.ListenAndServe() // Return error if any
}
