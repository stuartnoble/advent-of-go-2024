package dayone

import (
	"embed"
	"fmt"
	"sort"
	"strconv"
	"strings"

	"adventofgo/internal/utils"
)

//go:embed input.txt
var dataFiles embed.FS

func loadNumbers(fileBytes []byte) (numsLeft []int, numsRight []int, err error) {
	lines := strings.Split(string(fileBytes), "\n")

	// Assign cap to avoid resize on every append.
	numsLeft = make([]int, 0, len(lines))
	numsRight = make([]int, 0, len(lines))

	for _, l := range lines {
		// An empty line occurs at the end of the file when using Split.
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
		panic("numsLeft and numsRight must have the same length")
	}

	totalDistance := 0
	for i := range numsLeft {
		totalDistance += utils.Abs(numsRight[i] - numsLeft[i])
	}

	return totalDistance
}

// calculateSimilarityScore computes the similarity score based on the frequency
// of numbers in two slices.
func calculateSimilarityScore(numsLeft, numsRight []int) int {
	// Create a similarity map to track occurrences in numsLeft and numsRight.
	similarityMap := make(map[int][2]int)

	// Update the map for each item in numsLeft and numsRight.
	// Every the map entry is automatically allocated so we don't need to initialise
	// for the other value
	for i := range numsLeft {
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

func Solve() {
	fmt.Println("DAY 1")
	fmt.Println("-----")

	b, err := dataFiles.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	numsLeft, numsRight, err := loadNumbers(b)
	if err != nil {
		panic(err)
	}

	sort.Ints(numsLeft)
	sort.Ints(numsRight)

	totalDistance := calculateTotalDistance(numsLeft, numsRight)
	fmt.Println("ANSWER 1:", totalDistance)

	similarityScore := calculateSimilarityScore(numsLeft, numsRight)
	fmt.Println("ANSWER 2:", similarityScore)
}
