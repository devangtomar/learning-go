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

	responseSize("https://example.com")
}
