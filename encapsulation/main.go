package main

import (
	"errors"
	"fmt"
	"log"
	//"geo"
)

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d *Date) SetMonth(month int) error {
	if month > 12 || month < 1 {
		return errors.New("invalid month")
	}
	d.Month = month
	return nil
}

func (d *Date) SetYear(year int) {
	d.Year = year
}

func main() {
	date := Date{}
	date.SetYear(2019)
	fmt.Println(date.Year)

	// Adding validation to the setter methods
	err := date.SetMonth(14)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(date.Month)
}
