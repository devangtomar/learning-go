package main

import (
	"fmt"
)

func convertFeetToMeters(feet float64) float64{
	return feet*0.3048
}


func main() {
	var feet float64
	var meters float64
	fmt.Printf("Please enter the units in feets to convert to meters!\n")
	fmt.Scanf("%f", &feet)
	meters = convertFeetToMeters(feet)
	fmt.Printf("Converted meters from feet is %f", meters)
}