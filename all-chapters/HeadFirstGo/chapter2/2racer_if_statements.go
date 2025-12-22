/*Exercise 2: if Statements

We’re writing a program that asks the user for a racer’s name and the position they finished the race in, and prints out what type of medal they should get.

Complete the code sample below. The final call to fmt.Println at the bottom includes a variable that you’ll need to declare and assign a value to.
There are comments in the sample showing where to insert your code, and explaining what it should do.*/
package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter racer name: ")
	racer_name, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	racer_name = strings.TrimSpace(racer_name)

	fmt.Print("Enter racer rank: ")
	racer_rank_string, err := reader.ReadString('\n')
	if err != nil {
		log.Fatal(err)
	}
	racer_rank_string = strings.TrimSpace(racer_rank_string)
	racer_rank, err := strconv.ParseInt(racer_rank_string, 10, 64)
	if err != nil {
		log.Fatal(err)
	}

	// Write some code that looks at the
	// "rank" variable and sets the "medal" variable to an
	// appropriate string. If "rank" is 1, "medal" should
	// be set to "gold". If "rank" is 2, "medal" should be
	// "silver". A rank of 3 should get a "bronze" medal,
	// and any other rank should get a "participant" medal.
	var medal string
	switch racer_rank {
	case 1:
		medal = "gold"
		break
	case 2:
		medal = "silver"
		break
	case 3:
		medal = "bronze"
		break
	default:
		medal = "participant"
		break
	}

	fmt.Println(racer_name, "gets a", medal, "medal!")
}
