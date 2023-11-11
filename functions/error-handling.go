package main

import (
	"log"
	"fmt"
	"errors"
)

func main() {
	err := errors.New("Height can't be negative!")
	fmt.Println(err.Error())
	fmt.Println(err)
	log.Fatal(err)
}