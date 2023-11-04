// ! The rules for function names are the same as the rules for variable names:

// * A name must begin with a letter, followed by any number of
// * additional letters and numbers. (You’ll get a compile error if you
// * break this rule.)

// * Functions whose name begins with a capital letter are exported,
// * and can be used outside the current package. If you only need to
// * use a function inside the current package, you should start its name
// * with a lowercase letter.

// * Names with multiple words should use camelCase

package main

import (
	"fmt"
)

func sayHi() {
	fmt.Println("Hi!")
}

func main() {
	sayHi()
	repeatLine("hey there!", 10)
}

func repeatLine(line string, times int) {
	for i := 0; i < times; i++ {
		fmt.Println(line)
	}
}
