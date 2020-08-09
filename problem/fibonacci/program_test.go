package program_test

import (
	"testing"

	p "code-interview/problem/fibonacci"

	"github.com/stretchr/testify/require"
)

func TestFibonacci(t *testing.T) {
	tests := map[string]struct {
		input  int
		result int
	}{
		"1 should return 1": {
			input:  1,
			result: 1,
		},
		"2 should return 1": {
			input:  2,
			result: 1,
		},
		"3 should return 2": {
			input:  3,
			result: 2,
		},
		"5 should return 5": {
			input:  5,
			result: 5,
		},
		"6 should return 8": {
			input:  6,
			result: 8,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, test.result, p.Fibonacci(test.input))
		})
	}
}
