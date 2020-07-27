package program

// MatrixSpiral return matrix spiral
func MatrixSpiral(num int) [][]int {
	result := make([][]int, num)
	for i := range result {
		result[i] = make([]int, num)
	}

	count := 0
	row := 0
	column := num - 1

	for i := 0; i < num; i++ {
		count++
		result[row][i] = count
	}

	for j := 1; j < num; j++ {
		count++
		result[j][column] = count
	}

	for k := num - 2; k >= 0; k-- {
		count++
		result[num-1][k] = count
	}

	return result
}
