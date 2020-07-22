package palindrome

// IsPalindrome1 check if a word is a palindrome
func IsPalindrome1(word string) bool {
	newWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}
	return word == newWord
}

// IsPalindrome2 is optimized version of IsPalindrome2
func IsPalindrome2(word string) bool {
	for i := 0; i < len(word); i++ {
		ptr := len(word) - 1 - i
		if string(word[i]) == string(word[ptr]) {
			continue
		}
		return false
	}

	return true
}
