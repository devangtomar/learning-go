package main

import "fmt"

func main() {
	var arrayExample [4]string
	for i := range len(arrayExample) {
		arrayExample[i] = "hey"
	}
	fmt.Println(arrayExample)

	var notes [7]string
	notes[0] = "do"
	notes[1] = "re"
	notes[2] = "mi"
	fmt.Println(notes[0])
	fmt.Println(notes[1])
}
