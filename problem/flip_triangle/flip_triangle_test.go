package flipTriangle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	ft "code-interview/problem/flip_triangle"
)

func TestFlipTriangle(t *testing.T) {
	tests := map[string]struct {
		level  int
		result string
	}{
		"level1": {
			level:  1,
			result: "*",
		},
		"level3": {
			level:  3,
			result: "  *\n **\n***",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := ft.FlipTriangle(test.level)
			require.Equal(t, test.result, res)
		})
	}
}
