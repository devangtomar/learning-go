package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	files, err := os.ReadDir(".")
	if err != nil {
		fmt.Println("Error encountered!", err)
		return
	}
	fmt.Println("All the files in current dir : ", files)

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("Enter a number!")
	text, _ := reader.ReadString('\n')

	text = strings.TrimSuffix(text, "\n")

	fmt.Println("You entered the following string : ", text)
}
