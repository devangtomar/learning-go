package main

import "fmt"

func greeting(myChannel chan string) { // function taking argument as channel for values with a type string
	myChannel <- "hi" // Passing "hi" string into channel
}

func main() {
	myChannel := make(chan string)
	go greeting(myChannel)   // That how you call the function with a newly created channel
	// fmt.Println(<-myChannel) // Passing the incoming channel into print function!

	// We could also do this..
	recievedVar := <-myChannel // to assign to a variable!
	fmt.Println(recievedVar)
}
