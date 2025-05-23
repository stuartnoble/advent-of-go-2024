package daythree

import (
	"embed"
	"fmt"
	"regexp"
	"strconv"
)

type DayThreePuzzle struct {
	programCode string
}

type Factors struct {
	leftFactor  int
	rightFactor int
}

//go:embed daythree_input.txt
var fileSystem embed.FS

func extractFactors(programCode string) (factors []Factors, err error) {
	pattern := regexp.MustCompile(`mul\((\d*),(\d*)\)`)
	matches := pattern.FindAllStringSubmatch(programCode, -1)

	factors = make([]Factors, len(matches))
	for i, match := range matches {
		leftFactor, leftErr := strconv.Atoi(match[1])
		rightFactor, rightErr := strconv.Atoi(match[2])

		if leftErr != nil || rightErr != nil {
			return nil, err
		}

		factors[i] = Factors{
			leftFactor:  leftFactor,
			rightFactor: rightFactor,
		}
	}

	return factors, nil
}

func calculateSumOfProducts(factors []Factors) (sumOfProducts int) {
	for _, factor := range factors {
		product := factor.leftFactor * factor.rightFactor
		sumOfProducts = sumOfProducts + product
	}

	return sumOfProducts
}

func (p *DayThreePuzzle) LoadData() *DayThreePuzzle {
	data, err := fileSystem.ReadFile("daythree_input.txt")
	if err != nil {
		panic(err)
	}

	puzzleInput := string(data)
	p.programCode = puzzleInput

	return p
}

func (p *DayThreePuzzle) Solve() int {
	multiples, err := extractFactors(p.programCode)

	if err != nil {
		fmt.Println("Error parsing daythree_input.txt:", err)
		return 0
	}

	sumOfProducts := calculateSumOfProducts(multiples)

	return sumOfProducts
}
