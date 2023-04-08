package greetings

import (
	"errors"
	"fmt"
)

/* Function to welcome users */
func Hello(name string) (string, error) {
	//validation
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf("Hi, %v. Welcome user!", name)

	return message, nil
}
