package main

import (
	"fmt"
	"os"
)

func main() {
	// Slices declaration ~ Arrays without size limit 👀
	// var myArray [5]int -> array note the size
	// var mySlice [5]int -> slice no size specified
	var notes []string
	// Unlike with array variables, declaring a slice variable does’t automatically.. create a slice. For that, you can call the built-in make function.
	notes = make([]string, 7)
	// Once the slice is created, you assign and retrieve elements using the same syntax you would for an array
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[2])
	fmt.Println(notes[4])

	// Short declaration for slices
	primes := make([]int, 5)
	primes[0] = 3
	primes[1] = 6
	fmt.Println(primes[0])
	fmt.Println(len(primes))

	// Looping through a slice
	letters := []string{"a", "b", "c"}
	for _, letter := range letters {
		fmt.Println(letter)
	}

	// Another way of looping a slice!
	for i := 0; i < len(letters); i++ {
		fmt.Println(letters[i])
	}

	// Slice literals
	moreLetters := []string{"a", "b", "c"}
	fmt.Println(moreLetters)

	// Slice operators
	underlyingArray := [5]string{"a", "b", "c", "d", "e"}
	slice1 := underlyingArray[1:3] // Slice is nothing but an underlying array
	slice2 := underlyingArray[:3]
	fmt.Println(slice1)
	fmt.Println(slice2)
	underlyingArray[1] = "p"
	fmt.Println(slice2) // So changing the underlying array change the slice as well!

	// Adding elements into a slice
	slice := []string{"a", "b"}
	fmt.Println(slice, len(slice))
	slice = append(slice, "c")
	fmt.Println(slice, len(slice))

	// At default slice stores default value for that type! Bool ~> False, int ~> 0 etc.

	// Slice operator can be used on other slices
	underlyingAnotherArray := [5]string{"a", "b", "c", "d", "e"}
	slice5 := underlyingAnotherArray[1:]
	fmt.Println(slice5)

	// Command lines args
	fmt.Println(os.Args[2:]) // taking only after 1st input command lines args
}
