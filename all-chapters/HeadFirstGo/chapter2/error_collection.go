package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	fmt.Println("Enter your grade:")
	reader := bufio.NewReader(os.Stdin)
	input_grade, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Entered:", input_grade, " now getting the alpabet grade for you...")
	parsed_grade, err := strconv.ParseInt(input_grade, 10, 64)
	if err != nil {
		log.Fatal(err)
	}
	if parsed_grade >= 50 && parsed_grade <= 90 {
		fmt.Println("You got more than 50 but less than 90!")
	} else if parsed_grade >= 90 {
		fmt.Println("You got more than 90!")
	} else {
		fmt.Println("You got less than 50!")
	}
}
