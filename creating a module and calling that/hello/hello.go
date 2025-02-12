package main

import (
	"fmt"

	"example.com/greetings"
)

func main() {
	message := greetings.Hello("Pratham")
	fmt.Println(message)
}
