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

// KadaneAlgorithm2 optimized version of v1
// Time complexity -> O (n)
// Space complexity -> O(1)
func KadaneAlgorithm2(data []int) int {
	tmpMax := data[0]
	finalMax := data[0]

	for i := 1; i < len(data); i++ {
		tmpMax = max(data[i], tmpMax+data[i])
		finalMax = max(finalMax, tmpMax)
	}
	return finalMax
}
