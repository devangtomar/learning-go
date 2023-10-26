package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var status string
	s := "\t formerly surrounded by space \n"
	input := strings.TrimSpace(s) // this will get rid of the newline..
	// grade, err := strconv.ParseInt(input, 10, 64) // added base and bitSize arguments to ParseInt
	grade, err := strconv.ParseFloat(input, 64) // added base and bitSize arguments to ParseInt
	if err != nil {                             // corrected the error check
		fmt.Println("Some error occurred: ", err)
	}
	if grade >= 60 {
		status = "passing" // removed := to assign value to existing variable
	} else {
		status = "failing" // removed := to assign value to existing variable
	}
	fmt.Println(grade)
	fmt.Println("Your acquired grade is: ", status)
}
