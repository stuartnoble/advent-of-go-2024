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

	expectedSafeUnmodifiedReports := 2
	expectedSafeModifiedReports := 4

	dayTwoPuzzle := DayTwoPuzzle{}
	dayTwoPuzzle.reports = reports
	actualSafeUnmodifiedReports, actualSafeModifiedReports := dayTwoPuzzle.Solve()

	if actualSafeUnmodifiedReports != expectedSafeUnmodifiedReports {
		t.Errorf("Result from SolvePartOne() was incorrect; want: %v, got: %v.",
			actualSafeUnmodifiedReports,
			expectedSafeUnmodifiedReports)
	}

	if actualSafeModifiedReports != expectedSafeModifiedReports {
		t.Errorf("Result from SolvePartOne() was incorrect; want: %v, got: %v.",
			actualSafeModifiedReports,
			expectedSafeModifiedReports)
	}
}
