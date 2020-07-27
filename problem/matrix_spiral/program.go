package program

// MatrixSpiral return matrix spiral
func MatrixSpiral(num int) [][]int {
	result := make([][]int, num)
	for i := range result {
		result[i] = make([]int, num)
	}

	count := 0

	idxCol1, idxRow1 := 0, 0
	idxCol2, idxRow2 := num-1, num-1

	for idxRow1 <= idxRow2 && idxCol1 <= idxCol2 {
		for i := idxCol1; i <= idxCol2; i++ {
			count++
			result[idxRow1][i] = count
		}
		idxRow1++

		for i := idxRow1; i <= idxRow2; i++ {
			count++
			result[i][idxCol2] = count
		}
		idxCol2--

		for i := idxCol2; i >= idxCol1; i-- {
			count++
			result[idxRow2][i] = count
		}
		idxRow2--

		for i := idxRow2; i >= idxRow1; i-- {
			count++
			result[i][idxCol1] = count
		}
		idxCol1++
	}

	return result
}
