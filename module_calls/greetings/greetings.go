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

// HelloMultiple returns a map that would associate
// each named people with a greeting message
func HelloMultiple(names []string) (map[string]string, error) {
	// map to associate names with messages
	messages := make(map[string]string)

	// Loop through names, calling Hello func with each loop
	for _, name := range names {
		message, err := Hello(name)
		if err != nil {
			return nil, err
		}

		// associate retrieved message with the name
		messages[name] = message
	}

	return messages, nil
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
