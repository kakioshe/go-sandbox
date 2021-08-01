package greetings

import (
	"errors"
	"fmt"
	"math/rand"
	"time"
)

// Hello will return a greeting for `name` person.
func Hello(name string) (string, error) {
	// If no name was given, returns an error
	if name == "" {
		return "", errors.New("Empty Name")
	}

	// Return a greeting that embeds `name` in message.
	message := fmt.Sprintf(randomFormat(), name)
	return message, nil
}

// init sets initial values for variables used in the func
func init() {
	rand.Seed(time.Now().UnixNano())
}

// randomFormat returns a set of greeting message
// Returned message is selected at random
func randomFormat() string {
	formats := []string{
		"Hi, %v. Welcome!",
		"Great to see you , %v!",
		"Hail, %v! Well met!",
	}

	// returns a randomly selected message format
	// by specifying a random index for the slice of format
	return formats[rand.Intn(len(formats))]
}
