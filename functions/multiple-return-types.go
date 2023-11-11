package main

import "fmt"

func manyReturns() (int, bool, string) {
	return 1, false, "hello"
}

func paintNeeded(width float64, height float64) (float64, error) {
	if width < 0 {
		return 0, fmt.Errorf("a width of %0.2f is invalid", width)
	}
	if height < 0 {
		return 0, fmt.Errorf("a height of %0.2f is invalid", height)
	}
	area := height * width
	return area / 10.0, nil
}

func main() {
	fmt.Println(manyReturns())
	amount, err := paintNeeded(4.2, -3.0)
	fmt.Println(err)
	fmt.Printf("%0.2f litres needed\n", amount)
}
