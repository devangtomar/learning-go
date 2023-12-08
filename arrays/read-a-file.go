package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func readAFile() {
	file, err := os.Open("arrays/arrays.go")
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	err = file.Close()
	if err != nil {
		log.Fatal(err)
	}
	if scanner.Err() != nil {
		log.Fatal(scanner.Err())
	}
}

func readAFileIntoArray(fileName string) ([3]float64, error) {
	var numbers [3]float64
	file, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	scanner := bufio.NewScanner(file)
	i := 0
	for scanner.Scan() {
		numbers[i], err = strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			log.Fatal(err)
		}
		i += 1
	}
	err = file.Close()
	if err != nil {
		return numbers, err
	}

	if scanner.Err() != nil {
		return numbers, scanner.Err()
	}
	return numbers, nil

}

func main() {
	// readAFile()
	fileName := "./arrays/data.txt"
	var numbers_output [3]float64
	numbers_output, err := readAFileIntoArray(fileName)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(numbers_output)
}
