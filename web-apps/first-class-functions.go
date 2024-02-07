package main

import (
	"fmt"
	"log"
	"net/http"
)

func sayHi() {
	fmt.Println("Hi!")
}

func sayBye() {
	fmt.Println("Bye!")
}

func twice(theFunction func()) {
	theFunction()
	theFunction()
}

func main() {
	var myFunction func()
	myFunction = sayHi // First class function.. passing functions to other functions!
	myFunction()

	// Passing functions to other functions!
	twice(sayHi)
	twice(sayBye)

	// Functions as types!
	http.HandleFunc("/hello", sayHi)
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
