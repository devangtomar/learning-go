package main

import (
	"fmt"
)

func main() {
	paintNeeded(4.2, 3.0)
	paintNeeded(1.6, 1.0)
	paintNeeded(3.3, 8.0)
}

func paintNeeded(width float64, height float64) {
	area := width * height
	fmt.Printf("%.2f litres needed\n", area)
}