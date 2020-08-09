package program

// Fibonacci return fibonacci number based on index num
func Fibonacci(num int) int {
	data := []int{0, 1}

	for i := 2; i <= num; i++ {
		res := data[i-1] + data[i-2]
		data = append(data, res)
	}

	return data[num]
}
