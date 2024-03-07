package main

import "fmt"

func calmDown() {
	recover()
}

func freakOut() {
	defer calmDown()
	panic("oh no!")
}

func main() {

	// Recover function

	// Go offers a built-in recover function that can stop a program
	// from panicking. We’ll need to use it to exit the program gracefully

	fmt.Println(recover()) // it does nothing and returns nil!

	// If you call recover when a program is panicking, it will stop the panic. But
	// when you call panic in a function, that function stops executing. So there’s
	// no point calling recover in the same function as panic, because the panic
	// will continue anyway

	// But there's a way..
	freakOut()
}
