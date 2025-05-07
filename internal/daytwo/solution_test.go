package daytwo

import (
	"testing"
)

func TestSolve(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	expectedSafeReports := 2

	dayTwoPuzzle := DayTwoPuzzle{}
	dayTwoPuzzle.reports = reports
	actualSafeReports := dayTwoPuzzle.Solve()

	if actualSafeReports != expectedSafeReports {
		t.Errorf("Result from Solve() was incorrect; got: %v, want: %v.",
			expectedSafeReports,
			actualSafeReports)
	}
}
