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
