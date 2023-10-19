package main

import (
	"encoding/csv"
	"fmt"
	"os"
	_ "reflect"
)

func main() {
	print_csv()
}

func print_csv() {
	// Open the CSV file
	csv_file, err := os.OpenFile("./abc.csv", os.O_RDONLY, os.ModePerm)

	if err != nil {
		fmt.Println("File doesn't exists!")
		return
	}
	reader := csv.NewReader(csv_file)
	var i int = 0
	// fmt.Println(reflect.TypeOf(reader))
	for {
		// Read a record from the CSV file
		record, err := reader.Read()
		if err != nil {
			break
		}
		if i == 0 {
			// break
			goto next
			// fmt.Println(reflect.TypeOf(record))
		}
	next:
		// fmt.Println(record)
		for _, value := range record {
			// fmt.Println(index)
			fmt.Println(value)
		}
	}
}
