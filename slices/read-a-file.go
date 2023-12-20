package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func GetFloats(fileName string) ([]float64, error) {
	var numbers []float64
	file, err := os.Open("./slices.go")
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Println("Error occurred: ", err)
		}
		numbers = append(numbers, number)
	}
	err = file.Close()
	if err != nil {
		fmt.Println("Error occurred: ", err)
	}
	if scanner.Err() != nil {
		fmt.Println("Error occurred: ", err)
	}
	return numbers, nil
}
