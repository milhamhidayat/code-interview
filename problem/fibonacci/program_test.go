package program

import (
	"testing"

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
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			require.Equal(t, test.result, p.Fibonacci())
		})
	}
}
