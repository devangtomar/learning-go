// average2 conducts an average sum of numbers
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
)

func average(numbers ...float64 ) (float64) {
	var sum float64
	for _, num := range numbers {
		sum += num
	}
	return sum/float64(len(numbers))
}

func main() {
	args := os.Args[1:]
	var numbers []float64
	fmt.Printf("Calculating average of all the numbers that user input!")
	for _, arg := range args {
		number, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			log.Fatal(err)
		}
		numbers = append(numbers, number)
	}
	fmt.Printf("\nAverage of all the numbers you gave is %0.2f", average(numbers...))
}
