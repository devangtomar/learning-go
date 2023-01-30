/*
A method that is primarily used to set the value of an
encapsulated field is known as a setter method. Setter
methods often include validation logic, to ensure the new value
being provided is valid.

Since setter methods need to modify their receiver, their
receiver parameter should have a pointer type.

It’s conventional for setter method names to be in the form
SetX where X is the name of the field being set.

A method that is primarily used to get the value of an
encapsulated field is known as a getter method.

It’s conventional for getter method names to be in the form X
where X is the name of the field being set. Some other
programming languages favor the form GetX for getter method
names, but you should not use that form in Go.

Methods defined on an outer struct type live alongside methods
promoted from an embedded type.

An embedded type’s unexported methods don’t get promoted
to the outer type.
*/

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
