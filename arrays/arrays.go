package main

import (
	"fmt"
	"time"
)

func main() {

	// How to declare a array and iterate over a array
	var arrayExample [4]string
	for i := range len(arrayExample) {
		arrayExample[i] = "hey"
	}
	fmt.Println(arrayExample)

	// How to set a particular index
	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
	fmt.Println(notes[6]) // empty string

	// Integer type of array
	var primes [5]int
	primes[0] = 2
	primes[1] = 3
	fmt.Println(primes[0])
	fmt.Println(primes[3]) // should print 0 ~ since it's not assigned!

	// Array of time.Time values
	var dates [3]time.Time
	dates[0] = time.Unix(1257894000, 0)
	dates[1] = time.Unix(1447920000, 0)
	dates[2] = time.Unix(1557894000, 0)
	fmt.Println(dates[1])

	// Array literals

	var more_notes [7]string = [7]string{"do", "re", "mi", "do", "re", "mi", "do"}
	fmt.Println(more_notes[0])
	var more_primes [5]int = [5]int{2, 3, 5, 7, 11}
	fmt.Println(more_primes[2])

	// short variable declaration
	much_more_notes := [7]string{"do", "re", "mi", "do", "re", "mi", "do"}
	much_more_primes := [5]int{2, 3, 5, 7, 11}
	fmt.Println(much_more_notes[1])
	fmt.Println(much_more_primes[1])
}
