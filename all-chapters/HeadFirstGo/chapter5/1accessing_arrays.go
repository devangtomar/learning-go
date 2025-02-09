/*
Exercise 1: Accessing Arrays

Create an array with 5 string elements, holding English weekday names:
"Monday", "Tuesday", "Wednesday", "Thursday", "Friday".
Then print each array element along with its index.

You can assign array elements individually, or you can use an array literal.
You can access the elements individually, use a for loop to get each element index,
or use a for ... range loop to loop over the elements themselves.
Better yet, try all these techniques!
*/
package main

import (
	"fmt"
)

func main() {
	var week [5]string = [5]string{"Monday", "Tuesday", "Wednesday", "Thursday", "Friday"}
	for i, day := range week {
		fmt.Printf("Index for array: %d and Value for array: %s\n", i, day)
	}

	//OR
	var days [5]string = [5]string{} // "Monday", "Tuesday", "Wednesday", "Thursday", "Friday"
	days[0] = "Monday"
	days[1] = "Tuesday"
	days[2] = "Wednesday"
	days[3] = "Thursday"
	days[4] = "Friday"
	for j := 0; j < len(days); j++ {
		fmt.Printf("%d %s\n", j, days[j])
	}
}
