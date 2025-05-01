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

func sumList(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}

	// Recursive case: Calculate the sum of the first element and the rest of the list
	return numbers[0] + sumList(numbers[1:])
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

	sum := 0
	for i, num := range numsLeft {
		sum += numsRight[i] - num
	}

	distances := make([]int, 0, len(numsLeft))
	for index, numLeft := range numsLeft {
		distances = append(distances, Abs(numsRight[index]-numLeft))
	}

	totalDistance := sumList(distances)

	fmt.Println("ANSWER 1:", totalDistance)

	similarityMap := make(map[int][]int, len(numsLeft))

	for i := range numsLeft {
		numLeft := numsLeft[i]
		numRight := numsRight[i]

		if value, ok := similarityMap[numLeft]; ok {
			value[0]++
		} else {
			similarityMap[numLeft] = []int{1, 0}
		}

		if value, ok := similarityMap[numRight]; ok {
			value[1]++
		} else {
			similarityMap[numRight] = []int{0, 1}
		}
	}

	similarityScore := 0
	for key, value := range similarityMap {
		similarityScore += value[0] * value[1] * key
	}

	fmt.Println("ANSWER 2:", similarityScore)
}

func main() {
	fmt.Println("Advent of Code 2024")
	fmt.Println("-------------------")

	puzzle1()
}
