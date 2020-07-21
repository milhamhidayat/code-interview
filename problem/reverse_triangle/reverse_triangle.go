package reversetriangle

// ReverseTriangle will return reverse string with descending order level
func ReverseTriangle(level int) string {
	var result string

	for row := 1; row <= level; row++ {
		for col := level; col >= row; col-- {
			result += "*"
		}
		if row < level {
			result += "\n"
		}
	}

	return result
}
