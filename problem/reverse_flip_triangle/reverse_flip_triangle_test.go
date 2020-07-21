package reversefliptriangle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	rft "code-interview/problem/reverse_flip_triangle"
)

func TestReverseFlipTriangle(t *testing.T) {
	tests := map[string]struct {
		level  int
		result string
	}{
		"level1": {
			level:  1,
			result: "*",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := rft.ReverseFlipTriangle(test.level)
			require.Equal(t, test.result, res)
		})
	}
}
