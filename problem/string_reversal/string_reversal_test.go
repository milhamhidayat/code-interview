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
		"sonic": {
			word:   "sonic",
			result: "cinos",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := sr.StringReversal(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
