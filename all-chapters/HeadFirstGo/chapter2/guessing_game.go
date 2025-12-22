//Guessing game "Guess the number"
package main

/*
V1. Generate 1..100 and store it
V2. Offer to guess and store the guess
V3. If less - message, if more - message
V4. 10 attempts to guess, with a reminder
V5. if X = Y, output Success and stop replaying
V6. if attempts are over - sad message
task: -; 321; 23; 12; 13; nothing;
*/
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
	//1. generate the number
	scs := time.Now().Unix()
	rand.Seed(scs)
	fmt.Println("Chose a number from 1 to 100, can you guess?")
	seed := rand.Intn(100) + 1
	////////////////////////////////////

	//2. read from keyboard & 4. limit to 10 attempts
	reader := bufio.NewReader(os.Stdin)
	win := false
	// for guesses := 0; guesses < 10; guesses++ {
	for guesses := 0; guesses < 10; {
		fmt.Println("Remaining", 10-guesses, "attempts...")
		fmt.Print("Guess now:")
		input, err := reader.ReadString('\n') //read until Enter
		if err != nil {
			log.Fatal(err)
		}
		input = strings.TrimSpace(input)  //remove \n
		guess, err := strconv.Atoi(input) //string to number
		if err != nil {
			log.Fatal(err)
		}

		guesses ++
		////////////////////////////////////

		//3. Guess
		if guess < seed {
			fmt.Println("Oops, your guess is LESS.")
		} else if guess > seed {
			fmt.Println("Oops, your guess is MORE")
		} else { //5. Guess correct
			win = true
			fmt.Println("\n\tGood job! You guessed it! B-)")
			break
		}
		////////////////////////////////////
	}
	//6. Sad message if no success
	if !win {
		fmt.Println("Didn't guess, or attempts are over. I thought of", seed)
	}
	////////////////////////////////////
}
