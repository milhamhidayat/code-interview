package palindrome

// IsPalindrome check if a word is a palindrome
func IsPalindrome(word string) bool {
	newWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}
	return word == newWord
}
