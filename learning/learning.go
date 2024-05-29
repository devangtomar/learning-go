package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	students := map[string]int{"Devang": 2, "LTT": 3}
	fmt.Println("Hey there!")
	fmt.Println(students)

	for key, value := range students {
		fmt.Printf("Key : %s and value : %d\n", key, value)
	}
	os.Create("./some-file.txt")
	os.Remove("./some-file.txt")
	Reader := bufio.NewReader(os.Stdin)
	input, err := Reader.ReadString('\n')
	if err != nil {
		log.Fatalf("There are issues with the code %s", err.Error())
	}
	fmt.Printf("Your input was.. %s", input)
}
