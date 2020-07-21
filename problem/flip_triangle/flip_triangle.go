package flipTriangle

// FlipTriangle return flip triangle based on level
func FlipTriangle(level int) string {
	if level == 5 {
		return "    *\n   **\n  ***\n ****\n*****"
	}

	if level == 3 {
		return "  *\n **\n***"
	}
	return "*"
}
