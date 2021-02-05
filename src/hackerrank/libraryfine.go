package hackerrank

func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) int32 {
	if y1 > y2 {
		return 10000
	}

	if m1 > m2 && y1 == y2 {
		return (m1 - m2) * 500
	}

	if d1 > d2 && m1 == m2 && y1 == y2 {
		return (d1 - d2) * 15
	}

	return 0
}
