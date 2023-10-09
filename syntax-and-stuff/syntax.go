package main

import "fmt"

func main() {
	variables()
}

func variables() {
	// first way

	var num_1 int
	num_1 = 5
	fmt.Println(num_1)

	// second way

	var num_2 int
	num_2 = 5
	fmt.Println(num_2)

	// third way

	num_3 := 6
	fmt.Println(num_3)

}
