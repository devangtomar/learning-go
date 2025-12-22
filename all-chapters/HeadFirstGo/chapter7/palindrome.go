package main

import "fmt"

func isPalindrome(n int) bool {
	original := n
	reversed := 0

	for n > 0 {
		reversed = reversed*10 + n%10
		n /= 10
	}

	return reversed == original
}

func main() {
	var number int

	fmt.Print("Enter a number: ")
	if _, err := fmt.Scan(&number); err != nil {
		fmt.Println("Invalid input")
		return
	}

	fmt.Println(isPalindrome(number))
}
