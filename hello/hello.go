package main

import (
	"fmt"
	"github.com/renatoshk/go-learning-module/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Renato")
	fmt.Println(message)
}
