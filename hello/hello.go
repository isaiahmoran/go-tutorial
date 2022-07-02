package main

import (
	"fmt"
	"moran.dev/greetings"
)

func main() {
	// Get a greeting message and print it.
	message := greetings.Hello("Isaiah")
	fmt.Println(message)
}
