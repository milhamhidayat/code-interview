package reversetriangle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	rt "code-interview/problem/reverse_triangle"
)

func TestReverseTriangle(t *testing.T) {
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
			result: "***\n**\n*",
		},
		"level5": {
			level:  5,
			result: "*****\n****\n***\n**\n*",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := rt.ReverseTriangle(test.level)
			require.Equal(t, test.result, res)
		})
	}
}
