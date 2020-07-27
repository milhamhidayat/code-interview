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

// LongestPalindromicSubstring2 return longest palindrome sub string
func LongestPalindromicSubstring2(word string) string {
	longest := word[:1]

	for i := 1; i < len(word); i++ {
		tmpLongest := ""
		// odd palindrome
		odd := check(word, i-1, i+1)

		// even palindrome
		even := check(word, i-1, i)

		if len(odd) > len(even) {
			tmpLongest = odd
		} else {
			tmpLongest = even
		}

		if len(tmpLongest) > len(longest) {
			longest = tmpLongest
		}
	}

	return longest
}

func check(word string, start, end int) string {
	for start >= 0 && end < len(word) {
		if word[start] != word[end] {
			break
		}

		start--
		end++
	}

	return word[start+1 : end]
}
