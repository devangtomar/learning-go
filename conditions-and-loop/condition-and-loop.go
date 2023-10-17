package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	// Understanding how methods works..
	time_package()
	string_package()
	how_to_comment()

}

func time_package() {
	// this can also be written as .. year := time.Now().Year()
	var now time.Time = time.Now()
	var year int = now.Year() // `now.Year()` is a method call. `Year()` is a method of the `time.Time` type.  It is associated with the `now` variable, which is of type `time.Time`, and it returns the year of the time represented by `now`.

	fmt.Println(year)
}

func string_package() {
	broken := "G# r#cks!"

	// This can also be written as .. fixed := strings.NewReplacer("#", "o").Replace(broken)
	replacer := strings.NewReplacer("#", "o") // this is a function..`strings.NewReplacer("#", "o")` is a call to a function named `NewReplacer` from the `strings` package. This function returns a new `*strings.Replacer` based on the arguments provided.
	fixed := replacer.Replace(broken)         // this is a method since it returns a string type
	fmt.Println(fixed)

}

func how_to_comment() {
	var single_line string = "this is a sample" // this is a single lines comment
	fmt.Println(single_line)
	/*
	   var multi_line string
	   multi_line = "this is also a sample"
	*/
}
