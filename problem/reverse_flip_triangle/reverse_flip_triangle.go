package reversefliptriangle

// ReverseFlipTriangle will return reverse flip triangle with descending order level
func ReverseFlipTriangle(level int) string {
	var result string

	for row := 1; row <= level; row++ {
		for col := 1; col <= level; col++ {
			if col < row {
				result += " "
				continue
			}
			result += "*"
		}
		if row < level {
			result += "\n"
		}
	}
	return result
}
