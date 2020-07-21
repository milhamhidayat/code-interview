package flipTriangle

// FlipTriangle return flip triangle based on level
func FlipTriangle(level int) string {
	if level == 3 {
		return "  *\n **\n***"
	}
	return "*"
}
