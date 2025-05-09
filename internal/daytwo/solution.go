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

func checkIsSafeDelta(levels []int, maxDelta int) (bool, int) {
	firstDistance := levels[1] - levels[0]
	secondDistance := levels[2] - levels[1]

	isTrending := (firstDistance > 0 && secondDistance > 0) ||
		(firstDistance < 0 && secondDistance < 0)

	isWithinLimit := utils.Abs(firstDistance) <= maxDelta &&
		utils.Abs(secondDistance) <= maxDelta

	// If firstDistance is greater than secondDistance, put the second number in a variable, otherwise put the third number in a variable

	return isTrending && isWithinLimit
}

func isSafeReport(report []int) (bool, bool) {
	// Start by assuming the report is safe
	var isSafeReport = true
	///var isSafeWhenUnmodified bool = false
	//var isSafeWhenModified bool = false
	var isModified bool = false

	for i := 0; i < len(report)-2; {
		isSafeReport = checkIsSafeDelta(report[i:i+3], maxLevelDelta)

		i++

		if !isSafeReport && !isModified {
			// Don't do it this way, it will change the original slice!
			// modifiedSlice := append(report[:i+1], report[i+2:]...)

			// Don't forget to include index you want (+1)
			modifiedReport := append([]int{}, report[:i]...)
			modifiedReport = append(modifiedReport, report[i+1:]...)
			report = modifiedReport

			isModified = true

			// Recheck the same index
			i--

			//isSafeWhenModified = isSafeWhenModified && checkIsSafeDelta(modifiedReport[i:i+3], maxLevelDelta)
		}
	}

	return isSafeReport && !isModified, isSafeReport && isModified
}

func getSafeReportCount(p *DayTwoPuzzle) (int, int) {
	safeUnmodifiedReportsCount := 0
	safeModifiedReportsCount := 0

	for i := range p.reports {
		isSafeWhenUnmodified, isSafeWhenModified := isSafeReport(p.re	ports[i])

		if isSafeWhenUnmodified {
			safeUnmodifiedReportsCount++
		}

		if isSafeWhenModified {
			safeModifiedReportsCount++
		}
	}

	return safeUnmodifiedReportsCount, safeModifiedReportsCount
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

func (p *DayTwoPuzzle) Solve() (int, int) {
	safeUnmodifiedReportsCount, safeModifiedReportsCount := getSafeReportCount(p)

	return safeUnmodifiedReportsCount, safeUnmodifiedReportsCount + safeModifiedReportsCount
}
