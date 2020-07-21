package fliptriangle

// FlipTriangle return flip triangle based on level
func FlipTriangle(level int) string {
	var result string

	for row := 1; row <= level; row++ {
		for col := level; col >= 1; col-- {
			if col > row {
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
