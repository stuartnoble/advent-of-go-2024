package daytwo

import (
	"embed"
	"fmt"

	"adventofgo/internal/utils"
)

type DayTwoPuzzle struct {
	reports [][]int
}

//go:embed daytwo_input.txt
var fileSystem embed.FS

const maxLevelDelta int = 3

func checkIsSafeDelta(levels []int, maxDelta int) bool {
	firstDistance := levels[1] - levels[0]
	secondDistance := levels[2] - levels[1]

	isTrending := (firstDistance > 0 && secondDistance > 0) ||
		(firstDistance < 0 && secondDistance < 0)

	isWithinLimit := utils.Abs(firstDistance) <= maxDelta &&
		utils.Abs(secondDistance) <= maxDelta

	return isTrending && isWithinLimit
}

// I would love for this method not to mutate all the things
func isSafeReport(report []int, allowSingleLevelRemoval bool) bool {
	var isSafeReport bool = true
	var singleLevelRemoved = false

	for i := len(report); i > 2; i-- {
		isSafeReport = checkIsSafeDelta(report[i-3:i], maxLevelDelta)

		// If we're not allowing the caller to remove a bad
		// level, just exit and mark the report as bad
		if !isSafeReport && (!allowSingleLevelRemoval || singleLevelRemoved) {
			return false
		}

		if !isSafeReport && allowSingleLevelRemoval && !singleLevelRemoved {
			modifiedReport := append([]int{}, report[:i-2]...)
			modifiedReport = append(modifiedReport, report[i-1:]...)

			// Reassign to report. It is now shorter but because we're
			// searching backwards, the next index accounts for the
			// change in size
			report = modifiedReport

			singleLevelRemoved = true
		}
	}

	return isSafeReport
}

func getSafeReportCount(p *DayTwoPuzzle, allowSingleLevelRemoval bool) int {
	numSafeReports := 0

	for i := range p.reports {
		if isSafeReport(p.reports[i], allowSingleLevelRemoval) {
			numSafeReports++
		}
	}

	return numSafeReports
}

func (p *DayTwoPuzzle) LoadData() *DayTwoPuzzle {
	data, err := fileSystem.ReadFile("daytwo_input.txt")
	if err != nil {
		panic(err)
	}

	puzzleInput, err := utils.LoadMatrix(data)

	if err != nil {
		fmt.Println("Error loading daytwo_input.txt:", err)
		return p
	}

	p.reports = puzzleInput

	return p
}

func (p *DayTwoPuzzle) SolvePartOne() int {
	return getSafeReportCount(p, false)
}

func (p *DayTwoPuzzle) SolvePartTwo() int {
	return getSafeReportCount(p, true)
}
