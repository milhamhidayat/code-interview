package stringreversal

// StringReversal return reverse word
func StringReversal(word string) string {
	var result string

	for i := len(word) - 1; i >= 0; i-- {
		result += string(word[i])
	}

	return result
}
