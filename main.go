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

	partOneAnswer, partTwoAnswer := dayTwoPuzzle.
		LoadData().
		Solve()

	fmt.Println("Part 1:", partOneAnswer)
	fmt.Println("Part 2:", partTwoAnswer)
}
