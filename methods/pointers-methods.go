package main

import "fmt"

type Number int

func (n Number) Double() {
	n *= 2
}

func (n *Number) DoubleChange() {
	*n *= 2
}

func main() {
	number := Number(4)
	fmt.Println("Original value of number: ", number)
	number.Double()
	fmt.Println("Number after calling Double: ", number) // 😅 will return the same value..
	number.DoubleChange()

	// No need to use &number since..
	// if you call a method requiring a value
	// receiver, Go will automatically get the value at the pointer for you and pass
	// that to the method.

	fmt.Println("Number after calling DoubleChange: ", number) // This will change the value..

}
