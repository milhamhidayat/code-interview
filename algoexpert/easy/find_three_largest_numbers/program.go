package program

import (
	"math"
)

// FindThreeLargestNumbers returns sorted array with 3 maximum value
func FindThreeLargestNumbers(data []int) []int {
	result := []int{math.MinInt8, math.MinInt8, math.MinInt8}

	for _, v := range data {
		if result[2] < v {
			assign(2, result, v)
		} else if result[1] < v {
			assign(1, result, v)
		} else if result[0] < v {
			assign(0, result, v)
		}
	}

	return result
}

func assign(idx int, result []int, value int) {
	if idx == 2 {
		result[0] = result[1]
		result[1] = result[2]
		result[2] = value
	} else if idx == 1 {
		result[0] = result[1]
		result[1] = value
	} else if idx == 0 {
		result[0] = value
	}
}
