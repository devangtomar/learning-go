package main

import "fmt"

type Litres float64
type Gallons float64

// You can return a custom type as well from a function
func ToGallons(l Litres) Gallons {
	return Gallons(l)
}

// You can return a custom type as well from a function
func ToLitres(g Gallons) Litres {
	return Litres(g)
}

func main() {
	var carFuel Gallons
	var busFuel Litres
	carFuel = Gallons(10.0)
	//busFuel = Litres(10.0)
	busFuel = 10.0
	fmt.Println(carFuel, busFuel)
	fmt.Println(Litres(Gallons(11.0)))

	carFuel += ToGallons(Litres(40.0))
	fmt.Println(carFuel)
}
