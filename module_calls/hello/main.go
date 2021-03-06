package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {
	// Set properties of predefined Logger, including
	// Log entry prefix, and flag to disable printing
	// the time, source file, and line number
	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	// slice of names
	names := []string{
		"GladOS",
		"Lumine",
		"Var",
	}

	// Request greeting message
	message, err := greetings.HelloMultiple(names)

	// if error was returned, print it to console and
	// exit program
	if err != nil {
		log.Fatal(err)
	}

	// if no error returned, print returned message
	fmt.Println(message)
}
