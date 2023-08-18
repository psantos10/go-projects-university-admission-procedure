package main

import "fmt"

func main() {
	var score1, score2, score3 float64

	fmt.Scan(&score1)
	fmt.Scan(&score2)
	fmt.Scan(&score3)

	average := (score1 + score2 + score3) / 3
	fmt.Println(average)

	fmt.Println("Congratulations, you are accepted!")
}
