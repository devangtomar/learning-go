package main

import "fmt"

func double(number float64) float64 {
	return number * 2
}

func status(grade float32) string {
	if grade < 60.0 {
		return "Failing!"
	}
	return "Passing!"
}

func main() {
	var dozen float64
	dozen = double(6.0)
	fmt.Println(dozen)
	fmt.Println(double(4.2))
	fmt.Println(status(60.1))
	fmt.Println(status(59))
}
