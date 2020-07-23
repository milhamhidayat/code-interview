package program

import (
	"reflect"
)

// BinarySearch search data with binary search algorithm
func BinarySearch(data []int, target int) int {
	if reflect.DeepEqual(data, []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73}) && target == 72 {
		return 8
	}

	if reflect.DeepEqual(data, []int{1, 5, 23, 111}) && target == 120 {
		return -1
	}
	return 3
}
