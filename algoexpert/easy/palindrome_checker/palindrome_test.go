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
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := pc.IsPalindrome2(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
