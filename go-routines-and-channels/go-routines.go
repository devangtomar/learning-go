// Goroutines let your program work on several different tasks at once.
// You goroutines can coordinate their work using channels, which let them send data to each other
// and synchronize so that one goroutine doesn't get ahead of another.

// Goroutines let you take full advantage of computers with multiple processors, so that your programs run as fast as possible!

package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func responseSize(url string) {
	fmt.Println("Getting ", url)
	response, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(len(body)))
}

func main() {
	response, err := http.Get("https://example.com")
	if err != nil {
		log.Fatal(err)
	}
	defer response.Body.Close()
	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(body))

	// Takes too long and one after another!
	// responseSize("https://example.com/")
	// responseSize("https://golang.org/")
	// responseSize("https://golang.org/doc")

	// Concurrency allows a program to pause one task and work on other tasks.
	// If a program is written to support concurrency, then it may also support
	// parallelism: running tasks simultaneously. A computer with only one
	// processor can only run one task at a time. But most computers these days
	// have multiple processors (or one processor with multiple cores). Your
	// computer may divide concurrent tasks among different processors to run
	// them at the same time. (It’s rare to manage this directly; the operating
	// system usually handles it for you.

	go responseSize("https://example.com/")
	go responseSize("https://golang.org/doc")
	go responseSize("https://golang.org/")
	time.Sleep(5 * time.Second)

	// Go statements can't be used with return values
	// sum = go add(x, y) ~> this wouldn't work..

	// The only practical way to use a channel is to communicate from one goroutine to another goroutine.
	// 1. Create a channel
	// 2. Write a function that receives a channel as a parameter. We’ll run this function in a separate goroutine, and use it to send values over the channel.
	// 3. Receive the sent values in our original goroutine.

	// Not only do channels allow you to send values from one goroutine to another, they ensure the sending goroutine has sent the value before the receiving goroutine attempts to use it.

	var myChannel chan float64     // declare a variable to hold a channel
	myChannel = make(chan float64) // Actually create the channel

	myNewChannel := make(chan float64) // Create a channel and declare it..
	fmt.Println(myNewChannel)
	// Sending and receiving values with channels

	// To send a value on a channel, you use the <- operator (that’s a less-than
	// symbol followed by a dash). It looks like an arrow pointing from the value
	// you’re sending to the channel you’re sending it on.

	myChannel <- 3.14

	// fmt.Println(myChannel)

	// You also use the <- operator to receive values from a channel, but the
	// positioning is different: you place the arrow to the left of the channel you’re
	// receiving from. (It kind of looks like you’re pulling a value out of the channel.

	// <- myChannel // like this
}
