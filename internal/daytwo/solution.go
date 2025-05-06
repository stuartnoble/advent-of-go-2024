package daytwo

import (
	"adventofgo/internal/utils"
	"fmt"
)

func compareLevels(first, second int) (isIncrease, isDecrease bool, distance int) {
	return first < second, first > second, utils.Abs(first - second)
}

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
