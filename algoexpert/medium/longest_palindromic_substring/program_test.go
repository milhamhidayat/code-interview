package program_test

import (
	"testing"

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
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.LongestPalindromicSubstring(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
