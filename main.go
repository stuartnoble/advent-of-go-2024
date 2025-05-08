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

	dayTwoPartOne := dayTwoPuzzle.
		LoadData().
		SolvePartOne()
	fmt.Println("Answer 1:", dayTwoPartOne)

	dayTwoPartTwo := dayTwoPuzzle.
		LoadData().
		SolvePartTwo()
	fmt.Println("Answer 2:", dayTwoPartTwo)
}
