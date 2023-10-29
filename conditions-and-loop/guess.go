package main

import (
	"bufio"
	"fmt"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	target := rand.Intn(100) + 1
	fmt.Println("I've choose a random number b/w 1 and 100.")
	fmt.Println("Can you guess it?")
	fmt.Println(target)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Make a guess..")
	input, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	input = strings.TrimSpace(input)
	guess, err := strconv.Atoi(input)
	if err != nil {
		log.Fatal(err)
	}

	if guess < target {
		fmt.Println("Oops your guess was LOW!")
	} else if guess > target {
		fmt.Println("Oops your guess was LOW!")
	} else {
		fmt.Println("Wow your guess was accurate!")
	}
}
