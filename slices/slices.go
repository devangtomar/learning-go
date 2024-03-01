package main

import "fmt"

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
	slice1 := underlyingArray[1:3]
	fmt.Println(slice1)
}
