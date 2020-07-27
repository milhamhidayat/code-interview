package program_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	p "code-interview/problem/max_chars"
)

func TestMaxChars(t *testing.T) {
	tests := map[string]struct {
		word   string
		result string
	}{
		"abbcccccc should return c": {
			word:   "abbcccccc",
			result: "c",
		},
		"x should return x": {
			word:   "x",
			result: "x",
		},
		"apple11231231 should return 1": {
			word:   "apple11231231",
			result: "1",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.MaxChars(test.word)
			require.Equal(t, test.result, res)
		})
	}
}
