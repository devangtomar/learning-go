package main

import (
	"errors"
	"fmt"
	"math"
)

// task 5,6,false,$$$$,11,30,true,hahaha
func main() {
	fmt.Println("Hello World!")
	fmt.Printf("%v", "\na")      // a
	fmt.Printf("%#v", "\na")     // "\na"
	fmt.Printf("%f liters", 1.9) // 1.900000 liters
	// % width-number.width-decimal-part type-verb-format
	fmt.Printf("%5.3f", 11.1337) // 11.134
	sayHello()
	sayTimesX("man...", 5)
	fmt.Println(sayAndReturn("fisting costs $300"))
	fmt.Printf("%6.2f", circleArea(-9999999.9999999))
	err := errors.New("Error in the eye!")
	fmt.Println(err.Error())
	fmt.Println(err)
	outFloat, outString, outInteger := returns()
	fmt.Println("Float:", outFloat, "String:", outString, "Integer:", outInteger)
}

func sayHello() { // greeting function
	fmt.Println("PC says \"Greetings!\"")
}

func sayTimesX(say string, X int) { // function 'say X times'
	for i := 0; i < X; i++ {
		fmt.Println(say)
	}
}

func sayAndReturn(say string) string { // function 'say and return'
	fmt.Println(say, " and returning...")
	return say
}

func circleArea(rad float64) float64 {
	return math.Pow(rad, 2) * math.Pi
}

func returns() (float64, string, int) {
	return 13.37, "l33t", 2021
}
