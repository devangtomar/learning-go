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
}
