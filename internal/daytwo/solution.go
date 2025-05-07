package daytwo

import (
	"fmt"
)

func isSafeDelta(levels []int) {

}

// Much less naive implementation that uses the fact that we know
// how many levels are in each report to quickly assess the safety
func isSafeReport(report [5]int) bool {
	isSafeDelta(report[0:2])
	isSafeDelta(report[1:3])
	isSafeDelta(report[2:4])

	// isAllSafeLevels := true
	// isIncreasing := false
	// isDecreasing := false

	// for i := 0; i < len(report)-1; i++ {
	// 	isIncrease, isDecrease, distance := compareLevels(report[i], report[i+1])

	// 	isIncreasing = isIncreasing || isIncrease
	// 	isDecreasing = isDecreasing || isDecrease

	// 	if isIncreasing == isDecreasing || distance >= 4 {
	// 		isAllSafeLevels = false
	// 		break
	// 	}
	// }

	return isAllSafeLevels
}

func Solve(reports [][5]int) {
	fmt.Println("DAY 2")
	fmt.Println("-----")

	for i := range reports {
		fmt.Println(isSafeReport(reports[i]))
	}
}
