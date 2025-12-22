package main

import (
	"fmt"
)

const feetToMeter = 0.3048

func convertFeetToMeters(feet float64) float64 {
	return feet * feetToMeter
}

func main() {
	var feet float64

	fmt.Print("Enter feet: ")

	if _, err := fmt.Scan(&feet); err != nil {
		fmt.Println("Invalid input:", err)
		return
	}

	meters := convertFeetToMeters(feet)

	fmt.Printf("Meters: %.2f\n", meters)
}
