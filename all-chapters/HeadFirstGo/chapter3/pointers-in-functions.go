package main

import (
	"fmt"
)

func main() {
	amount := 6
	double(amount)
	fmt.Println(amount, "at address", &amount, "\tDid NOT change...")
	double2(&amount)
	fmt.Println(amount, "at address", &amount, "\tDid change!")
	fmt.Printf("And now the value is %v", *&amount)
}

func double(num int) {
	num *= 2
}
func double2(num *int) {
	*num *= 2
}
