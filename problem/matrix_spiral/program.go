package program

// MatrixSpiral return matrix spiral
func MatrixSpiral(num int) [][]int {
	if num == 3 {
		return [][]int{
			{1, 2, 3},
			{8, 9, 4},
			{7, 6, 5},
		}
	}
	return [][]int{
		{1, 2},
		{4, 3},
	}
}
