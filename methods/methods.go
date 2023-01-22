// Methods are nothing but classes in golang..
// So based on that implement that 👀

package main

import (
	"fmt"
	"go/types"
)

type custom types.Error

type MyType string

func (e custom) printError() {
	fmt.Println(e)
}

func (m MyType) sayHi() { // Go uses receiver parameters instead of 'self' and 'this' keywords
	fmt.Println("Hi from", m)
}

func (m MyType) MethodWithParameters(number int, flag bool) {
	fmt.Println(m)
	fmt.Println(number)
	fmt.Println(flag)
}

func (m MyType) WithReturn() int {
	return len(m)
}

func main() {
	value := MyType("a MyType value")
	value.sayHi()
	fmt.Println(value)
	anotherValue := MyType("Another value")
	anotherValue.sayHi()

	// A method is (pretty much) just like a function
	moreValue := MyType("MyType value")
	moreValue.MethodWithParameters(4, true)

	// More example
	fmt.Println(value.WithReturn())
}
