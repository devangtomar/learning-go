package main

import (
	"fmt"
)

// You can define your own types as well via "type"
type part struct {
	description string
	count       int
}

type subscriber struct {
	name   string
	rate   float64
	active bool
}

type car struct {
	name     string
	topSpeed float64
}

func showInfo(p part) {
	fmt.Println("Description : ", p.description)
	fmt.Println("Count : ", p.count)
}

func defaultSubscriiber(name string) subscriber {
	var s subscriber
	s.name = name
	s.rate = 5.44
	s.active = true
	return s
}

func applyDiscount(s *subscriber) {
	s.rate = 9.89
}

func printInfo(s subscriber) {
	fmt.Println("Name : ", s.name)
	fmt.Println("Rate : ", s.rate)
	fmt.Println("Active? : ", s.active)
}

func main() {
	// struct is when you need to store more than one type of data
	var myStruct struct {
		name string
		age  int
	}
	myStruct.name = "pie"
	myStruct.age = 2
	fmt.Printf("%#v\n", myStruct.age)
	fmt.Printf("%#v\n", myStruct.name)

	// You can define your own types as well via "type"
	var porsche car
	porsche.name = "Porsche 911 R"
	porsche.topSpeed = 323
	fmt.Println("Name : ", porsche.name)
	fmt.Println("Top speed : ", porsche.topSpeed)

	var bolts part
	bolts.description = "Hex bolts"
	bolts.count = 24
	fmt.Println("Description : ", bolts.description)
	fmt.Println("Count : ", bolts.count)

	// Using custom types with functions
	showInfo(bolts)

	subscriber1 := defaultSubscriiber("Devang tomar")
	subscriber1.rate = 4.99
	printInfo(subscriber1)

	// Modifying a struct using a function
	var s subscriber
	applyDiscount(&s)
	fmt.Println(s.rate)
}
