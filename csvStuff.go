package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func print_csv() {
	// Open the CSV file
	csv_file, err := os.OpenFile("./abc.csv", os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Println("File doesn't exists!")
		return
	}

	reader := csv.NewReader(csv_file)

	for {
		// Read a record from the CSV file
		record, err := reader.Read()
		if err != nil {
			break
		}
		fmt.Println(record)
	}
}
