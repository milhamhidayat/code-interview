package palindrome_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	pc "code-interview/algoexpert/easy/palindrome_checker"
)

func TestPalindrome1(t *testing.T) {
	tests := map[string]struct {
		word   string
		result bool
	}{
		"civic is a palindrome": {
			word:   "civic",
			result: true,
		},
		"level is a palindrome": {
			word:   "level",
			result: true,
		},
		"example is not a palindrome": {
			word:   "example",
			result: false,
		},
		"a is a palindrome": {
			word:   "a",
			result: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := pc.IsPalindrome1(test.word)
			require.Equal(t, test.result, res)
		})
	}
}

func TestPalindrome2(t *testing.T) {
	tests := map[string]struct {
		word   string
		result bool
	}{
		"civic is a palindrome": {
			word:   "civic",
			result: true,
		},
		"level is a palindrome": {
			word:   "level",
			result: true,
		},
		"example is not a palindrome": {
			word:   "example",
			result: false,
		},
		"a is a palindrome": {
			word:   "a",
			result: true,
		},
		"noon is a palindrome": {
			word:   "noon",
			result: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := pc.IsPalindrome2(test.word)
			require.Equal(t, test.result, res)
		})
	}
}

func TestPalindrome3(t *testing.T) {
	tests := map[string]struct {
		word   string
		result bool
	}{
		"civic is a palindrome": {
			word:   "civic",
			result: true,
		},
		"level is a palindrome": {
			word:   "level",
			result: true,
		},
		"example is not a palindrome": {
			word:   "example",
			result: false,
		},
		"a is a palindrome": {
			word:   "a",
			result: true,
		},
		"noon is a palindrome": {
			word:   "noon",
			result: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := pc.IsPalindrome3(test.word)
			require.Equal(t, test.result, res)
		})
	}
}

func TestPalindrome4(t *testing.T) {
	tests := map[string]struct {
		word   string
		result bool
	}{
		"civic is a palindrome": {
			word:   "civic",
			result: true,
		},
		"level is a palindrome": {
			word:   "level",
			result: true,
		},
		"example is not a palindrome": {
			word:   "example",
			result: false,
		},
		"a is a palindrome": {
			word:   "a",
			result: true,
		},
		"noon is a palindrome": {
			word:   "noon",
			result: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := pc.IsPalindrome4(test.word)
			require.Equal(t, test.result, res)
		})
	}
}

func BenchmarkPalindrome1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.IsPalindrome1("saippuakivikauppias")
	}
}

func BenchmarkPalindrome2(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.IsPalindrome2("saippuakivikauppias")
	}
}

func BenchmarkPalindrome3(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.IsPalindrome3("saippuakivikauppias")
	}
}

func BenchmarkPalindrome4(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pc.IsPalindrome4("saippuakivikauppias")
	}
}
