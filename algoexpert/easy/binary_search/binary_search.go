package program

// BinarySearch search data with binary search algorithm
func BinarySearch(data []int, target int) int {
	leftIdx := 0
	rightIdx := len(data) - 1

	for leftIdx <= rightIdx {
		middleIdx := (leftIdx + rightIdx) / 2
		if target == data[middleIdx] {
			return middleIdx
		}
		if target < data[middleIdx] {
			rightIdx = middleIdx - 1
		} else if target > data[middleIdx] {
			leftIdx = middleIdx + 1
		}
	}

	return -1
}
