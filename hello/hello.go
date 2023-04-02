package main

import (
	"fmt"
	"github.com/renatoshk/go-learning-module/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)
	// Get a greeting message and print it.
	message, err := greetings.Hello("Renato")

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
