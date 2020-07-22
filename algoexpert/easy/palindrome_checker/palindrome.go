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
// O(n) time -> karena loop dari awal string sampai akhir string
// O(1) space -> karena simpan data hanya untuk variable ptr
func IsPalindrome2(word string) bool {
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
