package main

import (
	"fmt"
	"reflect"
)

func createPointer() *float64 {
	var myFloat = 98.5
	return &myFloat
}

func double(number int) {
	number *= 2
	number *= 2
}

func pointer() {
	amount := 6
	fmt.Println(amount)
	fmt.Println(&amount) // retrieves the variable's address!
}

func printPointer(myBoolPointer *bool) {
	fmt.Println(*myBoolPointer)
}

func main() {
	amount := 6
	double(amount)
	fmt.Println(amount)
	pointer()

	// pointer types..
	var myInt int
	fmt.Println((reflect.TypeOf(&myInt)))
	var myFloat float64
	fmt.Println((reflect.TypeOf(&myFloat)))
	var myBool bool
	fmt.Println((reflect.TypeOf(&myBool)))

	// Declaring a pointer pointing to a int type ~> myNewInt integer variable!
	var myNewInt int
	var myIntPointer *int
	myIntPointer = &myNewInt
	fmt.Println(myIntPointer)

	// Short declaration for a pointer variable!
	var myNewBool bool = false
	myBoolPointer := &myNewBool
	fmt.Println(myBoolPointer)

	// For printing value of a pointer!
	fmt.Println(*myBoolPointer)

	// Using pointers with functions

	var myFloatPointer *float64 = createPointer()
	fmt.Println(*myFloatPointer)

	var myFuncBool bool = true
	printPointer(&myFuncBool)

	// Fixing our “double” function using pointers ~> this will now change the original value of amt variable!
	amt := 7
	doublePointerFix(&amt)
	fmt.Println(amt)
}

func doublePointerFix(number *int) {
	*number *= 2
}
