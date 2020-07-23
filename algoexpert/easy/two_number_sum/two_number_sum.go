package program

import (
	"sort"
)

// TwoNumberSum return slice with value sum up to the target sum
// Time Complexity : O(n ^ 2), Space complexity: O(n), n is a number of  input
func TwoNumberSum(data []int, target int) []int {
	result := []int{}

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			sum := data[i] + data[j]
			if sum == target {
				result = append(result, data[i], data[j])
			}
		}
	}
	if len(result) > 1 {
		sort.Sort(sort.IntSlice(result))
	}
	return result
}

// TwoNumberSum2 return slice with value sum up to the target sum
// Time Complexity : O(n), Space complexity: O(n) , n is a number of  input
func TwoNumberSum2(data []int, target int) []int {
	num := map[int]int{}

	for _, v := range data {
		v2 := target - v
		_, ok := num[v]
		if ok {
			result := []int{v2, v}
			sort.Sort(sort.IntSlice(result))
			return result
		}

		num[v2] = v
	}

	return []int{}
}
