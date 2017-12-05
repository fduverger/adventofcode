package corruptionchecksum

import (
	"sort"
)

// CorruptionChecksum determine the difference between the largest
// value and the smallest value
func CorruptionChecksum(spreedsheet [][]int) int {
	checksum := 0
	for _, row := range spreedsheet {
		sort.Ints(row)
		checksum += row[len(row)-1] - row[0]
	}

	return checksum
}
