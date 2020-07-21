package triangle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tr "code-interview/problem/triangle1"
)

func TestTriangle(t *testing.T) {
	tests := map[string]struct {
		level  int
		result string
	}{
		"level1": {
			level:  1,
			result: `*`,
		},
		"level2": {
			level:  2,
			result: "* \n**",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := tr.Triangle(test.level)
			require.Equal(t, test.result, res)
		})
	}
}
