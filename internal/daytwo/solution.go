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

func isSafeReport(report []int) bool {
	// Start by assuming the report is safe
	var isSafeReport = true

	for i := 0; i < len(report)-2; i++ {
		isSafeReport = isSafeReport && checkIsSafeDelta(report[i:i+3], maxLevelDelta)
	}

	return isSafeReport
}

func isSafeReportWhenModified(report []int, indexToRemove int) bool {
	if indexToRemove == len(report) {
		return false
	}

	modifiedReport := append([]int{}, report[:indexToRemove]...)

	if indexToRemove < len(report)-1 {
		modifiedReport = append(modifiedReport, report[indexToRemove+1:]...)
	}

	if !isSafeReport(modifiedReport) {
		return isSafeReportWhenModified(report, indexToRemove+1)
	}

	return true
}

func getSafeReportCount(p *DayTwoPuzzle) (int, int) {
	safeUnmodifiedReportsCount := 0
	safeModifiedReportsCount := 0

	for i := range p.reports {
		isSafeReportResult := isSafeReport(p.reports[i])

		if isSafeReportResult {
			safeUnmodifiedReportsCount++
		}

		if !isSafeReportResult && isSafeReportWhenModified(p.reports[i], 0) {
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
