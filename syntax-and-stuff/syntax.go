package main

import (
	"fmt"
	"reflect"
	_ "reflect"
)

func main() {
	variables()
	playingWithTypes()
}

func variables() {
	// first way

	var num_1 int
	num_1 = 5
	fmt.Println(num_1)

	// second way

	var num_2 int
	num_2 = 5
	fmt.Println(num_2)

	// third way

	num_3 := 6
	fmt.Println(num_3)

}

func playingWithTypes() {
	var abc int = 6
	var name string = "devang"
	var marks float64 = 5.12
	var char rune = 'j'
	var decision bool

	fmt.Println(reflect.TypeOf(abc))
	fmt.Println(reflect.TypeOf(float64(abc)))
	fmt.Println(reflect.TypeOf(name))
	fmt.Println(reflect.TypeOf(marks))
	fmt.Println(reflect.TypeOf(char))
	fmt.Println(reflect.TypeOf(decision))
	fmt.Println(decision)
}
