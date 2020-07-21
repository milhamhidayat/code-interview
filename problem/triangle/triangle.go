package triangle

import "fmt"

// Triangle is a function to return triangle with n level
func Triangle(level int) string {
	var result string
	for row := 1; row <= level; row++ {
		for col := 1; col <= row; col++ {
			result += "*"
		}
		if row < level {
			result += "\n"
		}
	}
	return result
}

// Print will print the triangle based on level
func Print(level int) {
	for row := 1; row <= level; row++ {
		for col := 1; col <= row; col++ {
			if col <= row {
				fmt.Print("*")
			}
		}
		fmt.Println()
	}
}
