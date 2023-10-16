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
}

func time_package() {
	// this can also be written as .. year := time.Now().Year()
	var now time.Time = time.Now()
	var year int = now.Year()

	fmt.Println(year)
}

func string_package() {
	broken := "G# r#cks!"

	// This can also be written as .. fixed := strings.NewReplacer("#", "o").Replace(broken)
	replacer := strings.NewReplacer("#", "o")
	fixed := replacer.Replace(broken)
	fmt.Println(fixed)

}
