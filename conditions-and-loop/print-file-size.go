package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	fileInfo, err := os.Stat("./condition-and-loop.go")
	if err != nil {
		log.Fatal("Error encounterred here: ", err)
	}
	fmt.Println(fileInfo.Size() / 1024)
}
