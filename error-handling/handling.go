package main

import (
	"fmt"
	"log"
)

func MoreSocialize() error {
	defer fmt.Println("Goodbye!")
	fmt.Println("Hello!")
	return fmt.Errorf("I don't want to talk!")
	fmt.Println("Nice Weather, eh?")
	return nil
}

func Socialize() {
	/*
		The “defer” keyword ensures a function call takes place, even if the
		calling function exits early.

		Example : Ensuring files get closed using deferred function calls
	*/
	defer fmt.Println("Goodbye!") // This will call won't be run until all the remaining code in the socialise function runs..
	fmt.Println("Hello!")
	fmt.Println("Nice Weather, eh?")
}

func main() {
	// Three main things to know..
	// 1. Defer: How to defer cleanup actions so they happen even when there's an error
	// 2. Panic
	// 3. Recover

	Socialize()
	err := MoreSocialize()
	if err != nil {
		log.Fatal(err)
	}
}
