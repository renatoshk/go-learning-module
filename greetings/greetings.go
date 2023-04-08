package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

/* Function to welcome users */
func Hello(name string) (string, error) {
	//validation
	if name == "" {
		return "", errors.New("Empty name")
	}

	message := fmt.Sprintf(randomFormat(), name)

	return message, nil
}

//function to get random index
func init() {
	rand.Seed(time.Now().UnixNano())
}

//function to create random greetings
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome",
		"Pershendetje, %v. Welcome",
	}

	return formats[rand.Intn(len(formats))]
}
