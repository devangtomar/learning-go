package main

import (
	"fmt"
)

func main() {
	//  for keyword initialization statement (optional ~ not really needed); condition expression (required ~ mandatory); post statement (optional ~ not really needed) {
	// 	loop body..
	// }
	for x := 1; x < 10; x += 1 {
		fmt.Println("Iterating over a for loop.. iteration number", x)
	}

	for x := 1; true; x += 1 { // this is a infinite loop
		fmt.Println("Iterating over a for loop.. iteration number", x)
	}

	for x := 1; false; x += 1 { // this loop will never run
		fmt.Println("Iterating over a for loop.. iteration number", x)
	}
}
