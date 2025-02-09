package main

import (
	"fmt"
)

func checkLength(length float64) (float64, error) {
	if length < 0 {
		return 0, fmt.Errorf("Length %0.2f is impossible", length)
	}
	lng := length * 10
	return lng, nil
}

func main() {
	leng, err := checkLength(-1)
	fmt.Println(leng, err)
}
