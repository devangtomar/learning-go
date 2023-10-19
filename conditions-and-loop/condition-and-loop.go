package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// Understanding how methods works..
	time_package()
	string_package()
	how_to_comment()
	input_output()
	error_checking()

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

func input_output() {
	fmt.Println("Enter a grade: ")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n') // we are ignoring the error returned from the ReadString method..
	fmt.Print(input)

}

func error_checking() {
	bool, err := strconv.ParseBool("true") // returns an error if the string can't be converted -> boolean
	fmt.Print(err)
	fmt.Print(bool)
	file, err := os.Open("file.txt") // returns an error if the file can't be opened
	fmt.Print(file)
	response, err := http.Get("http://golang.org") // returns an error if the page can't be accessed
	fmt.Print(response)

}
