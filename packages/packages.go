package main

import (
	"fmt"
	"log"
	"github.com/headfirstgo/keyboard"
)

func main() {
	fmt.Print("Enter a grade: ")
	grade, err := keyboard.GetFloat()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(grade)
}