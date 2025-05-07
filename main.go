package main

import (
	"fmt"

	"adventofgo/internal/dayone"
	"adventofgo/internal/daytwo"
)

func main() {
	fmt.Println("Advent of Go 2024")
	fmt.Println("===================")

	dayone.Solve()

	fmt.Println("DAY 2")
	fmt.Println("-----")

	dayTwoPuzzle := daytwo.DayTwoPuzzle{}

	dayTwoAnswerOne := dayTwoPuzzle.
		LoadData().
		Solve()
	fmt.Println("Answer 1:", dayTwoAnswerOne)
}
