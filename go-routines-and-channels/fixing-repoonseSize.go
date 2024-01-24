package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	URL  string
	Size int
}

func responseSize(url string, channel chan int) {
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
	// fmt.Println(string(len(body)))
	channel <- len(body)
}

func responseSizeWithStruct(url string, channel chan Page) {
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
	// fmt.Println(string(len(body)))
	channel <- Page{URL: url, Size: len(body)}
}

func main() {
	sizes := make(chan int)
	go responseSize("https://example.com/", sizes)
	go responseSize("https://golang.org/doc", sizes)
	go responseSize("https://golang.org/", sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)
	fmt.Println(<-sizes)

	// Make this even more faster!
	urls := []string{"https://example.com/", "https://golang.org/doc", "https://golang.org/"}
	for _, url := range urls {
		go responseSize(url, sizes)
	}
	for i := 0; i < len(urls); i++ {
		fmt.Println(<-sizes)
	}

	// Fetching and updating our channel to carry a struct for more defined output
	pages := make(chan Page)
	for _, url := range urls {
		go responseSizeWithStruct(url, pages)
	}
	for i := 0; i < len(urls); i++ {
		page := <-pages
		fmt.Printf("%s: %d\n", page.URL, page.Size)
	}

}
