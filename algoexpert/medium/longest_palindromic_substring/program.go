package program

// LongestPalindromicSubstring1 return longest palindrome sub string
func LongestPalindromicSubstring1(word string) string {
	longest := ""
	for i := 0; i < len(word); i++ {
		for j := i + 1; j <= len(word); j++ {
			newWord := string(word[i:j])
			result := isPalindrome(newWord)
			if result && len(newWord) > len(longest) {
				longest = newWord
			}
		}
	}
	return longest
}

func isPalindrome(word string) bool {
	for i := 0; i < len(word); i++ {
		ptr := len(word) - 1 - i
		if i == ptr {
			return true
		}
		if word[i] != word[ptr] {
			return false
		}
	}

	return true
}
