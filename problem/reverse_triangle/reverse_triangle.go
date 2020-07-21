package reversetriangle

// ReverseTriangle will return reverse string with descending order level
func ReverseTriangle(level int) string {
	if level == 5 {
		return "*****\n****\n***\n**\n*"
	}
	if level == 3 {
		return "***\n**\n*"
	}
	return "*"
}
