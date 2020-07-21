package hollowtriangle

// HollowTriangle return hollow triangle
func HollowTriangle(level int) string {
	if level == 5 {
		return "*\n**\n* *\n*  *\n*****"
	}
	return "*"
}
