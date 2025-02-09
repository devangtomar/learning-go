package main

import "fmt"

func scoreSummary(name string, pts1 float64, pts2 float64, pts3 float64) {
	var avg float64 = (pts1 + pts2 + pts3) / 3
	fmt.Printf("%10s\t| %8.2f\t| %8.2f\t| %8.2f\t| %8.2f", name, pts1, pts2, pts3, avg)
	fmt.Println()
}

func main() {
	fmt.Printf("%10s | %8s | %8s | %8s | %8s\n",
		"Name", "Russian", "Math", "ICT", "Average")
	for i := 0; i < 35; i++ {
		fmt.Print("-")
	}
	fmt.Println("------------------------------------------------------")
	scoreSummary("Sergey", 95.4, 82.3, 74.6)
	scoreSummary("Kirill", 79.3, 99.1, 82.5)
	scoreSummary("Anton", 82.2, 95.4, 77.6)
}

// % width-number.width-decimal point-type-format
