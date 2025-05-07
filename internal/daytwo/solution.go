package daytwo

import (
	"adventofgo/internal/utils"
	"fmt"
)

type DayTwoPuzzle struct {
	reports [][]int
}

const maxLevelDelta int = 3

func isSafeDelta(levels []int, maxDelta int) bool {
	firstDistance := levels[1] - levels[0]
	secondDistance := levels[2] - levels[1]

	isTrending := (firstDistance > 0 && secondDistance > 0) ||
		(firstDistance < 0 && secondDistance < 0)

	isWithinLimit := utils.Abs(firstDistance) <= maxDelta &&
		utils.Abs(secondDistance) <= maxDelta

	return isTrending && isWithinLimit
}

func isSafeReport(report []int) bool {
	var isSafeReport bool = true

	for i := 0; i <= len(report)-3; i++ {
		isSafeReport = isSafeReport && isSafeDelta(report[i:i+3], maxLevelDelta)
	}

	return isSafeReport
}

func (p *DayTwoPuzzle) LoadData() *DayTwoPuzzle {
	puzzleInput, err := utils.LoadMatrix("input.txt")

	if err != nil {
		fmt.Println("Error loading daytwo/input.txt:", err)
		return p
	}

	p.reports = puzzleInput

	return p
}

func (p *DayTwoPuzzle) Solve() int {
	numSafeReports := 0

	for i := range p.reports {
		if isSafeReport(p.reports[i]) {
			numSafeReports++
		}
	}

	return numSafeReports
}
