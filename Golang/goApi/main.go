package main

import "fmt"

func main() {
	server := NewAPIServer(":8080") // Corrected function name
	if err := server.Run(); err != nil {
		fmt.Println("Error starting server:", err) // Handle error
	}
}
