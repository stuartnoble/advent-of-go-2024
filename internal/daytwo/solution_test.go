package daytwo

import (
	"testing"
)

func TestSolvePartOne(t *testing.T) {
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
	actualSafeReports, _ := dayTwoPuzzle.Solve()

	if actualSafeReports != expectedSafeReports {
		t.Errorf("Result from SolvePartOne() was incorrect; want: %v, got: %v.",
			expectedSafeReports,
			actualSafeReports)
	}
}

func TestSolvePartTwo(t *testing.T) {
	reports := [][]int{
		{7, 6, 4, 2, 1},
		{1, 2, 7, 8, 9},
		{9, 7, 6, 2, 1},
		{1, 3, 2, 4, 5},
		{8, 6, 4, 4, 1},
		{1, 3, 6, 7, 9},
	}

	expectedSafeReports := 4

	dayTwoPuzzle := DayTwoPuzzle{}
	dayTwoPuzzle.reports = reports
	_, actualSafeReports := dayTwoPuzzle.Solve()

	if actualSafeReports != expectedSafeReports {
		t.Errorf("Result from SolvePartTwo() was incorrect; want: %v, got: %v.",
			expectedSafeReports,
			actualSafeReports)
	}
}
