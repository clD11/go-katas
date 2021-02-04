package hackerrank

import (
	"fmt"
)

const baseDays = 243

func DayOfTheProgrammer(year int32) string {
	day := baseDays

	if isLeapYear(year) {
		day += 1
	}
	if year == 1918 {
		day -= 13
	}

	return fmt.Sprintf("%d.09.%d", 256-day, year)
}

func isLeapYear(year int32) bool {
	if year <= 1917 && year%4 == 0 {
		return true
	}
	if year%400 == 0 || (year%4 == 0 && year%100 != 0) {
		return true
	}
	return false
}
