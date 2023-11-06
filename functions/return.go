package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}

func main() {
	var dozen float64
	dozen = double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
}
