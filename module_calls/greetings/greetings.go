package greetings

import "fmt"

// Hello will return a greeting for `name` person.
func Hello(name string) string {
	// Return a greeting that embeds `name` in message.
	message := fmt.Sprintf("Hi, %v. Welcome!", name)
	return message
}
