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

/* Function to welcome multiple users */
func Hellos(names []string) (map[string]string, error) {
	messages := make(map[string]string)

	for _, name := range names {
		message, err := Hello(name)

		if err != nil {
			return nil, err
		}
		messages[name] = message
	}

	return messages, nil
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
