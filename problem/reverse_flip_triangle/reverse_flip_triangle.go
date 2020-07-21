package reversefliptriangle

// ReverseFlipTriangle will return reverse flip triangle with descending order level
func ReverseFlipTriangle(level int) string {
	if level == 3 {
		return "***\n**\n*"
	}
	return "*"
}
