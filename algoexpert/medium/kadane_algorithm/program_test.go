package program_test

import (
	"testing"

	p "code-interview/algoexpert/medium/kadane_algorithm"

	"github.com/stretchr/testify/require"
)

func TestKadaneAlgorithm(t *testing.T) {
	tests := map[string]struct {
		data   []int
		result int
	}{
		"[1] should return 1": {
			data:   []int{1},
			result: 1,
		},
		"[-2, 1] should return 1": {
			data:   []int{-2, 1},
			result: 1,
		},
		"[1, 2, 3] should return 6": {
			data:   []int{1, 2, 3},
			result: 6,
		},
		"[-1,2,3] should return 5": {
			data:   []int{-1, 2, 3},
			result: 5,
		},
		"[3, 5, -9, 4, 7] should return 11": {
			data:   []int{3, 5, -9, 4, 7},
			result: 11,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.KadaneAlgorithm(test.data)
			require.Equal(t, test.result, res)
		})
	}
}
