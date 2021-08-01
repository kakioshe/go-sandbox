package greetings

import (
	"errors"
	"fmt"
)

// Hello will return a greeting for `name` person.
func Hello(name string) (string, error) {
	// If no name was given, returns an error
	if name == "" {
		return "", errors.New("Empty Name")
	}

	// Return a greeting that embeds `name` in message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message, nil
}
