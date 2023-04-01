package greetings

import "fmt"

/* Function to welcome users */
func Hello(name string) string {
	message := fmt.Sprintf("Hi, %v. Welcome user!", name)
	return message
}
