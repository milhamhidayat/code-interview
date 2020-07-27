package palindrome

// IsPalindrome1 check if a word is a palindrome
func IsPalindrome1(word string) bool {
	newWord := ""
	for i := len(word) - 1; i >= 0; i-- {
		newWord += string(word[i])
	}
	return word == newWord
}

// IsPalindrome2 is optimized version of IsPalindrome1
// O(n) time -> karena loop dari awal string sampai akhir string
// O(1) space -> karena simpan data hanya untuk variable ptr
func IsPalindrome2(word string) bool {
	length := len(word) - 1
	for i := 0; i <= length; i++ {
		ptr := length - i
		if word[i] != word[ptr] {
			return false
		}
		if i >= ptr {
			return true
		}
	}

	return true
}

// IsPalindrome3 optimized version of IsPalindrome2
// O(n/2) -> O(n) time -> karena loop dari awal string sampai akhir string
// O(4) -> O(n) space -> karena simpan data hanya untuk variable ptr
func IsPalindrome3(word string) bool {
	mid := len(word) / 2
	left := mid - 1
	right := 0
	if len(word)%2 == 0 {
		right = mid
	} else {
		right = mid + 1
	}

	for left >= 0 && right < len(word) {
		if word[left] != word[right] {
			return false
		}

		left--
		right++
	}

	return true
}

// IsPalindrome4 is currently the most optimized version of palindrome checker
func IsPalindrome4(word string) bool {
	front := 0
	back := len(word) - 1

	for back > front {
		if word[front] != word[back] {
			return false
		}

		front++
		back--
	}

	return true
}
