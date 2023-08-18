package main

import "fmt"

const threshold = 60.0

func main() {
	var score1, score2, score3 float64

	fmt.Scan(&score1)
	fmt.Scan(&score2)
	fmt.Scan(&score3)

	average := (score1 + score2 + score3) / 3
	fmt.Println(average)

	if average >= threshold {
		fmt.Println("Congratulations, you are accepted!")
	} else {
		fmt.Println("We regret to inform you that we will not be able to offer you admission.")
	}
}
