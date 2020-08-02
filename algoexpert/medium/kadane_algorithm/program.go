package program

import (
	"math"
)

// KadaneAlgorithm will return maximum of sub array value
// Time complexity -> O(n ^ 2)
// Space complexity -> O(1)
func KadaneAlgorithm(data []int) int {
	res := math.MinInt8

	for i := 0; i < len(data); i++ {
		tmpRes := data[i]
		for j := i + 1; j < len(data); j++ {
			tmpRes = tmpRes + data[j]
		}
		res = max(res, tmpRes)
	}

	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
