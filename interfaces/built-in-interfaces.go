package main

import "fmt"

type ComedyError string
type Gallons float64
type Litres float64
type Millilitres float64

func (c ComedyError) Error() string {
	return string(c)
}

// The empty interface doesn’t require any methods to satisfy it, and so it’s satisfied by all types.
type Anything interface{}

func main() {
	// Error interface
	err := fmt.Errorf("A height of %0.2f is invalid", -2.33333)
	fmt.Println(err.Error())
	fmt.Println(err)

	moreErr := ComedyError("What's a programmer's favorite beef? Logger!")
	fmt.Println(moreErr)

	// Stringer interface
	fmt.Println(Gallons(12.9090900))
	fmt.Println(Litres(12.9090900))
	fmt.Println(Millilitres(12.9090900))

	fmt.Printf("%0.2f gal\n", Gallons(12.9090900))
	fmt.Printf("%0.2f lit\n", Litres(12.9090900))
	fmt.Printf("%0.2f ml\n", Millilitres(12.9090900))

}
