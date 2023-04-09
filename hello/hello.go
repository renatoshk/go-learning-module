package main

import (
	"fmt"
	"github.com/renatoshk/go-learning-module/greetings"
	"log"
)

func main() {
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	names := []string{"Naruto", "Sasuke", "Itachi"}
	// Get a greeting message and print it.
	message, err := greetings.Hellos(names)

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(message)
}
