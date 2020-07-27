package program_test

import (
	"testing"

	p "code-interview/algoexpert/medium/longest_palindromic_substring"

	"github.com/stretchr/testify/require"
)

func TestLongestPalindromicSubstring(t *testing.T) {
	tests := map[string]struct {
		word   string
		result string
	}{
		"cbbd should return bb": {
			word:   "cbbd",
			result: "bb",
		},
		"abaxyzzyxf should return xyzzyx": {
			word:   "abaxyzzyxf",
			result: "xyzzyx",
		},
		"a should return a": {
			word:   "a",
			result: "a",
		},
		"bbc should return bb": {
			word:   "bbc",
			result: "bb",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.LongestPalindromicSubstring1(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
