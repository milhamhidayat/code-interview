package palindrome_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	pc "code-interview/algoexpert/easy/palindrome_checker"
)

func TestPalindrome(t *testing.T) {
	tests := map[string]struct {
		word   string
		result bool
	}{
		"civic is a palindrome": {
			word:   "civic",
			result: true,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := pc.IsPalindrome(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
