package hollowtriangle

// HollowTriangle return hollow triangle
func HollowTriangle(level int) string {
	var result string

	for row := 1; row <= level; row++ {
		for col := 1; col <= row; col++ {
			if col == 1 {
				result += "*"
			} else if row == level {
				result += "*"
			} else if col == row {
				result += "*"
			} else {
				result += " "
			}
		}
		if row < level {
			result += "\n"
		}
	}
	return result
}
