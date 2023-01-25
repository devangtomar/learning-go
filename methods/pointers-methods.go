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
	fmt.Println("Number after calling DoubleChange: ", number) // This will change the value..

}
