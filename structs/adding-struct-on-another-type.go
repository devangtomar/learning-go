package main

import (
	"fmt"
	"github.com/headfirstgo/magazine"
)

//type Subscriber struct {
//	Name        string
//	Rate        float64
//	Active      bool
//	HomeAddress Address
//}
//
//type Employee struct {
//	Name        string
//	Salary      float64
//	HomeAddress Address
//}
//
//type Address struct {
//	Location string
//}

func main() {
	fmt.Println("Hey there!")
	//address := magazine.Address{Street: "123 Oak st", City: "Omaha", State: "NE", PostalCode: "545454"}
	subscriber := magazine.Subscriber{Name: "Aman singh"}
	subscriber.Address.Street = "123 oak Street" // This way you can set up a struct within another struct
}
