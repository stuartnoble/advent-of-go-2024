package main

import (
	"adventofgo/internal/daythree"
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
	twoOneAnswer, twoTwoAnswer := dayTwoPuzzle.
		LoadData().
		Solve()

	fmt.Println("Part 1:", twoOneAnswer)
	fmt.Println("Part 2:", twoTwoAnswer)

	fmt.Println("DAY 3")
	fmt.Println("-----")

	dayThreePuzzle := daythree.DayThreePuzzle{}
	threeOneAnswer := dayThreePuzzle.
		LoadData().
		Solve()

	fmt.Println("Part 1:", threeOneAnswer)
}
