package stringreversal_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	sr "code-interview/problem/string_reversal"
)

func TestStringReversal(t *testing.T) {
	tests := map[string]struct {
		word   string
		result string
	}{
		"convert sonic to cinos": {
			word:   "sonic",
			result: "cinos",
		},
		"convert amazon to nozama": {
			word:   "amazon",
			result: "nozama",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := sr.StringReversal(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
