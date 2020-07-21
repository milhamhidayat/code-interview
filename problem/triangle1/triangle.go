package triangle

// Triangle is a function to return triangle with n level
func Triangle(level int) string {
	var result string
	for row := 1; row <= level; row++ {
		result += "*"
	}
	return result
}

// only  print the triangle
func triangle(level int) {

}
