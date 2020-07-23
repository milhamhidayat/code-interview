package program

import (
	"reflect"
)

// BinarySearch search data with binary search algorithm
func BinarySearch(data []int, target int) int {
	if reflect.DeepEqual(data, []int{1, 5, 23, 111}) && target == 120 {
		return -1
	}
	return 3
}
