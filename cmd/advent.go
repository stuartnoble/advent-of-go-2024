package main

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

//go:embed data/day1.txt
var dataFiles embed.FS

func loadNumbers(fileBytes []byte) (numsLeft []int, numsRight []int, err error) {
	lines := strings.Split(string(fileBytes), "\n")

	// Assign cap to avoid resize on every append.
	numsLeft = make([]int, 0, len(lines))
	numsRight = make([]int, 0, len(lines))

	for _, l := range lines {
		// Empty line occurs at the end of the file when we use Split.
		if len(l) == 0 {
			continue
		}

		splitNums := strings.Split(l, "   ")

		// Use Atoi when we know it's an integer.
		// NOTE: Scanf is a more general approach.
		numLeft, errLeft := strconv.Atoi(splitNums[0])
		numRight, errRight := strconv.Atoi(splitNums[1])

		if errLeft != nil || errRight != nil {
			return nil, nil, err
		}

		numsLeft = append(numsLeft, numLeft)
		numsRight = append(numsRight, numRight)
	}

	return numsLeft, numsRight, nil
}

func calculateTotalDistance(numsLeft, numsRight []int) int {
	if len(numsLeft) != len(numsRight) {
		panic("slices must have the same length")
	}

	totalDistance := 0
	for i := 0; i < len(numsLeft); i++ {
		totalDistance += Abs(numsRight[i] - numsLeft[i])
	}

	return totalDistance
}

// calculateSimilarityScore computes the similarity score based on the frequency
// of numbers in two slices.
func calculateSimilarityScore(numsLeft, numsRight []int) int {
	// Create a similarity map to track occurrences in numsLeft and numsRight.
	similarityMap := make(map[int][2]int)

	// Update the map for numsLeft and numsRight.
	for i := 0; i < len(numsLeft); i++ {
		similarityMap[numsLeft[i]] = [2]int{
			similarityMap[numsLeft[i]][0] + 1,
			similarityMap[numsLeft[i]][1],
		}
		similarityMap[numsRight[i]] = [2]int{
			similarityMap[numsRight[i]][0],
			similarityMap[numsRight[i]][1] + 1,
		}
	}

	// Compute the similarity score.
	similarityScore := 0
	for key, value := range similarityMap {
		similarityScore += value[0] * value[1] * key
	}

	return similarityScore
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func puzzle1() {
	fmt.Println("DAY 1")

	b, err := dataFiles.ReadFile("data/day1.txt")
	if err != nil {
		panic(err)
	}

	numsLeft, numsRight, err := loadNumbers(b)
	if err != nil {
		panic(err)
	}

	// Part 1
	sort.Ints(numsLeft)
	sort.Ints(numsRight)

	totalDistance := calculateTotalDistance(numsLeft, numsRight)
	fmt.Println("ANSWER 1:", totalDistance)

	similarityScore := calculateSimilarityScore(numsLeft, numsRight)
	fmt.Println("ANSWER 2:", similarityScore)
}

func main() {
	fmt.Println("Advent of Code 2024")
	fmt.Println("-------------------")

	puzzle1()
}
