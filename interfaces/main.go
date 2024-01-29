package main

import (
	"fmt"
	"github.com/headfirstgo/gadget"
	"mypkg"
)

func playList(device gadget.TapePlayer, songs []string) {
	for _, song := range songs {
		device.Play(song)
	}
	device.Stop()
}

// Interface is a set of methods that certain values are expected to have.
type myInterface interface {
	methodWithoutParameters()      // Method name
	methodWithParameters(float64)  // Type of params ~> float64
	methodWithReturnValue() string // Type of return value
}

func main() {
	player := gadget.TapeRecorder{}
	mixtape := []string{"Jessie's Girl", "Whip it", "9 to 5"}
	fmt.Println(player)
	fmt.Println(mixtape)
	// This will result in an error.. ⬇️
	//playList(player, mixtape)
	/*
		The TapeRecorder type defines all the
		methods that the playList function needs, but we’re being blocked from
		using it because playList only accepts TapePlayer values
	*/

	var value mypkg.MyInterface
	value = mypkg.MyType(5)
	value.MethodWithoutParameters()
	value.MethodWithParameters(124.3)
	fmt.Println(value.MethodWithReturnValue())

	// You can only call methods defined as part of the interface
}
