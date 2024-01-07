package main

import (
	"fmt"
	"reflect"
)

// A variadic function is one that can be called with a varying
// number of arguments. To make a function variadic, use an ellipsis (...)
// before the type of the last (or only) function parameter in the function
// declaration.

func severalInts(numbers ...int) {
	fmt.Println(numbers)
}

func severalStrings(strings ...string) {
	fmt.Println(strings)
	fmt.Println(reflect.TypeOf(strings))
}

func average(numbers ...float64) float64 {
	var sum float64 = 0
	for _, number := range numbers {
		sum += number
	}
	return sum / float64(len(numbers))
}
func main() {
	severalStrings("a", "d", "f") // returns a slice
	severalInts(23, 2, 9)
	severalInts(23, 29)
	severalStrings("a", "d", "f", "f", "w")
	severalStrings("a", "d", "f", "f", "w")
	severalStrings()
	fmt.Println(average(100, 40))
	fmt.Println(average(100, 40, 34, 90))
	fmt.Println(average(1.00, 4.0, 34.9, 90.12))
}
