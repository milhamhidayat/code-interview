package triangle_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	tr "code-interview/problem/triangle"
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
			result: "*\n**",
		},
		"level3": {
			level:  3,
			result: "*\n**\n***",
		},
		"level4": {
			level:  4,
			result: "*\n**\n***\n****",
		},
		"level5": {
			level:  5,
			result: "*\n**\n***\n****\n*****",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := tr.Triangle(test.level)
			require.Equal(t, test.result, res)
		})
	}
}

func TestPrint(t *testing.T) {
	tests := map[string]struct {
		level int
	}{
		"level1": {
			level: 1,
		},
		"level2": {
			level: 2,
		},
		"level3": {
			level: 3,
		},
		"level4": {
			level: 4,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			tr.Print(test.level)
		})
	}
}
