package program_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	p "code-interview/algoexpert/easy/two_number_sum"
)

func TestTwoNumberSum(t *testing.T) {
	tests := map[string]struct {
		data   []int
		target int
		result []int
	}{
		"[4,6] target 10, return [4,6]": {
			data:   []int{4, 6},
			target: 10,
			result: []int{4, 6},
		},
		"[4,6,1,-3] target 3 return []": {
			data:   []int{4, 6, 1, -3},
			target: 3,
			result: []int{},
		},
		"[14] target 15 return []": {
			data:   []int{14},
			target: 15,
			result: []int{},
		},
		"[15] target 15 return [15]": {
			data:   []int{15},
			target: 15,
			result: []int{},
		},
		"[3,5,-4,8,11,1,-1,6] target 10 return [-1,11]": {
			data:   []int{3, 5, -4, 8, 11, 1, -1, 6},
			target: 10,
			result: []int{-1, 11},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.TwoNumberSum(test.data, test.target)
			require.Equal(t, test.result, res)
		})
	}
}
