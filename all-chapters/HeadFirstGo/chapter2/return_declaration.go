package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter something:")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	fmt.Println("You entered:", input)
}
