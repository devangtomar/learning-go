package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func input_output(input *bufio.Reader) string {

	fmt.Print("Enter a number!")
	text, _ := input.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")

	return text
}

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error encountered!", err)
		return
	}
	fmt.Println("All the files in current dir : ", files)
	reader := bufio.NewReader(os.Stdin)
	output := input_output(reader)

	fmt.Println("You entered the following string : ", output)
}
