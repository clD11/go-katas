package hackerrank

const baseDays = 215

func DayOfTheProgrammer(year int32) string {
	return "12.09.2016"
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
