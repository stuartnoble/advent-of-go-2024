package daytwo

type DayTwoPuzzle struct {
	reports [][5]int
}

const maxLevelDelta int = 3

func isSafeDelta(levels []int, maxDelta int) bool {
	firstDistance := levels[1] - levels[0]
	secondDistance := levels[2] - levels[1]

	isTrending := (firstDistance > 0 && secondDistance > 0) ||
		(firstDistance < 0 && secondDistance < 0)

	isWithinLimit := firstDistance <= maxDelta && secondDistance <= maxDelta

	return isTrending || isWithinLimit
}

// Much less naive implementation that uses the fact that we know
// how many levels are in each report to quickly assess the safety
func isSafeReport(report [5]int) bool {
	return isSafeDelta(report[0:3], maxLevelDelta) &&
		isSafeDelta(report[1:3], maxLevelDelta) &&
		isSafeDelta(report[2:3], maxLevelDelta)
}

func (p *DayTwoPuzzle) LoadData() *DayTwoPuzzle {
	p.reports = nil

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
