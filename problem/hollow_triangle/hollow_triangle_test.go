package hollowtriangle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	ht "code-interview/problem/hollow_triangle"
)

func TestHollowTriangle(t *testing.T) {
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
			res := ht.HollowTriangle(test.level)
			require.Equal(t, test.result, res)
		})
	}
}
