package main

import (
	"log"
	"net/http"
)

func viewHandler(writer http.ResponseWriter, request *http.Request) {
	message := []byte("Hello web!")
	_, err := writer.Write(message)
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/hello", viewHandler)
	log.Println("Server is running on http://localhost:8080")
	err := http.ListenAndServe("localhost:8080", nil)
	log.Fatal(err)
}
