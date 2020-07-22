package program

import "sort"

// TwoNumberSum return slice with value sum up to the target sum
func TwoNumberSum(data []int, target int) []int {
	result := []int{}

	if len(data) == 1 {
		return result
	}

	for i := 0; i < len(data); i++ {
		for j := i + 1; j < len(data); j++ {
			sum := data[i] + data[j]
			if sum == target {
				result = append(result, data[i], data[j])
			}
		}
	}
	sort.Sort(sort.IntSlice(result))
	return result
}
