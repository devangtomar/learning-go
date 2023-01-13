package main

import "fmt"

func main() {
	//var myMap map[string]float64 // That's how you declare a map in golang..
	// With key as string and values as float64 with "map keyword"

	var ranks map[string]int     // declaring a map variable
	ranks = make(map[string]int) // actually create the map
	ranks["Dragon"] = 2

	ranksMetal := make(map[string]int) // create a map and declare a variable to hold it
	ranksMetal["gold"] = 1
	ranksMetal["silver"] = 2
	ranksMetal["bronze"] = 3
	fmt.Println(ranksMetal["bronze"])

	// Map literals

	myMap := map[string]int{"a": 1, "b": 6}
	fmt.Println(myMap)
	fmt.Println(myMap["c"])      // will print "0" value
	emptyMap := map[string]int{} // Creating a empty map!
	fmt.Println(emptyMap)

	// Checking if key value pairs exist or not
	var value int
	var ok bool
	value, ok = myMap["c"]
	fmt.Println(value, ok)

	// For removing key/value pairs via delete function
	delete(myMap, "b")
	fmt.Println(myMap)

	// Looping a map in golang
	for key, value := range myMap {
		fmt.Println(key, value)
	}
}
