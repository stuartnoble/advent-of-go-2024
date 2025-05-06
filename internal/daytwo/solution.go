package daytwo

import (
	"fmt"

	"adventofgo/internal/utils"
)

func compareLevels(first, second int) (isIncrease, isDecrease bool, distance int) {
	return first < second, first > second, utils.Abs(first - second)
}

// Possible recursive method: pass in previous distance, first, second
// Initial call: 0, 7, 6 -> safe
// 2nd call: -1, 6, 4 -> safe
// 3rd call: -2, 4, 2 -> safe

// Initial call: 0, 1, 2 -> safe
// 2nd call: 1, 2, 7 -> unsafe (+5)

// Initial call: 0, 1, 3 -> safe
// 2nd call: 2, 3, 2 -> unsafe (+2 then -1)
func isSafeReport(report [5]int) bool {
	isAllSafeLevels := true
	isIncreasing := false
	isDecreasing := false

	for i := 0; i < len(report)-1; i++ {
		isIncrease, isDecrease, distance := compareLevels(report[i], report[i+1])

		isIncreasing = isIncreasing || isIncrease
		isDecreasing = isDecreasing || isDecrease

		if isIncreasing == isDecreasing || distance >= 4 {
			isAllSafeLevels = false
			break
		}
	}

	return isAllSafeLevels
}

func Solve(reports [][5]int) {
	fmt.Println("DAY 2")
	fmt.Println("-----")

	for i := range reports {
		fmt.Println(isSafeReport(reports[i]))
	}
}
