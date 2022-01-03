package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	// Set properties of the predefined Logger, including
	// the log entry prefix and a flag to disable printing
	// the time, source file, and line number.
	log.SetPrefix("greetings: ")
	log.SetFlags(log.LUTC | log.Ldate)

	// Get a greeting message and print it.
	message, err := greetings.Hello("Waqas")

	// If an error was returned, print it to the console and exit the program
	if err != nil {
		log.Fatal(err)
	}
	log.Print("This is a log statement")
	// If no error was returned, print the returned message to the console
	fmt.Println(message)

	// A slice of names.
	names := []string{"Waqas", "Altaf", "Taslim"}
	// Request greeting messages for the names.
	messages, err := greetings.Hellos(names)
	if err != nil {
		log.Fatal(err)
	}
	// If no error was returned, print the returned map of
	// messages to the console.
	fmt.Println(messages)
}

//go mod edit -replace example.com/greetings=../greetings
