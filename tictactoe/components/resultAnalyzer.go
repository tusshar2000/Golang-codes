package components

func ResultAnalyzer(b [9]int) int {
	sums := [8]int{0, 0, 0, 0, 0, 0, 0, 0}

	sums[0] = b[2] + b[4] + b[6]
	sums[1] = b[0] + b[3] + b[6]
	sums[2] = b[1] + b[4] + b[7]
	sums[3] = b[2] + b[5] + b[8]
	sums[4] = b[0] + b[4] + b[8]
	sums[5] = b[6] + b[7] + b[8]
	sums[6] = b[3] + b[4] + b[5]
	sums[7] = b[0] + b[1] + b[2]

	for _, v := range sums {
		if v == 3 {
			return 1
		} else if v == 300 {
			return 2
		}
	}
	return 0
}
