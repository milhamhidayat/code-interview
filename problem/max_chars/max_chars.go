package program

// MaxChars return maximum character of a word
func MaxChars(word string) string {
	var (
		result  rune
		maxChar int
	)
	chars := make(map[rune]int)

	for _, l := range word {
		v, ok := chars[l]
		if !ok {
			chars[l] = 1
			maxChar = 1
			result = l
			continue
		}
		total := v + 1
		chars[l] = total
		if total > maxChar {
			result = l
			maxChar = total
		}
	}

	return string(result)
}
