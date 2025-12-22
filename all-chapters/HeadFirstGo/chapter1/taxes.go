// Common format verbs
// Verb	Meaning
// %s	string
// %d	integer
// %f	float
// %.2f	float with 2 decimals
// %t	boolean
// %c	character
// %v	default value
// %T	type of value

package main

import (
	"fmt"
)

func main() {
	var price int = 100
	var currency string = "Rs"
	fmt.Println("Price is", price, " ", currency)
	fmt.Printf("Price is %v %s", price, currency)

	var taxRate float64 = 0.08
	var tax float64 = float64(price) * taxRate
	fmt.Println("Taxes amount to", tax, "R.")
	fmt.Printf("Taxes amount to %f Rs", tax)

	var total float64 = float64(price) + tax
	fmt.Println("Total cost is", total, "R.")
	fmt.Printf("Total cost is %v %s", total, currency)

	var availableFunds int = 120
	fmt.Println(availableFunds, "R. available.")
	fmt.Printf("Available funds %d %s", availableFunds)
	fmt.Println("Within budget?", total <= float64(availableFunds))
}
