package program

import (
	"sort"
)

// ThreeNumSum return triplet slice value equal with target
// Time Complexity O (n ^ 3)
// Space complexity O (1)
func ThreeNumSum(data []int, target int) [][]int {
	sort.Ints(data)
	result := [][]int{}

	for i := 0; i < len(data)-1; i++ {
		for j := i + 1; j < len(data); j++ {
			for k := j + 1; k < len(data); k++ {
				v1 := data[i]
				v2 := data[j]
				v3 := data[k]

				if target == v1+v2+v3 {
					result = append(result, []int{v1, v2, v3})
				}
			}
		}
	}

	return result
}

// ThreeNumSum2 is optimized version of ThreeNumSum1
// Time Complexity O(n ^ 2) -> loop 2 kali
// Space Complexity O (n) -> worst case insert semua value array
func ThreeNumSum2(data []int, target int) [][]int {
	sort.Ints(data)

	result := [][]int{}
	lenData := len(data)

	for i := 0; i < len(data)-1; i++ {
		left := i + 1
		right := lenData - 1

		for left < right {
			sum := data[i] + data[left] + data[right]

			if sum == target {
				result = append(result, []int{data[i], data[left], data[right]})
				left++
				right--
			} else if sum > target {
				right--
			} else if sum < target {
				left++
			}
		}
	}
	return result
}
