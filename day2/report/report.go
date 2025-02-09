package report

import (
	"example.com/level"
)

func Safe(report []int) bool {
	var increasing bool

	for i := range report {
		// The report is safe when this for loop reaches the last element without returning
		if i+1 == len(report) {
			break
		}

		level1 := report[i]
		level2 := report[i+1]

		if i == 0 {
			increasing = level1 < level2
		}

		safe := level.Safe(level1, level2, increasing)

		if !safe {
			return false
		}
	}
	return true
}

func TolerantSafe(report []int) bool {
	if Safe(report) {
		return true
	} else {
		for i := 0; i < len(report); i++ {
			var partialReport []int

			for o := range report {
				if o != i {
					partialReport = append(partialReport, report[o])
				}
			}

			if Safe(partialReport) {
				return true
			}
		}
	}
	return false
}
