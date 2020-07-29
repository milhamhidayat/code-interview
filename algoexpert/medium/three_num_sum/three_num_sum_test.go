package program_test

import (
	"testing"

	p "code-interview/algoexpert/medium/three_num_sum"

	"github.com/stretchr/testify/require"
)

func TestThreeNumSum(t *testing.T) {
	tests := map[string]struct {
		data   []int
		target int
		result [][]int
	}{
		"[12, 3, 1, 2, -6, 5, -8, 6] target 6 should return triplet": {
			data:   []int{12, 3, 1, 2, -6, 5, -8, 6},
			target: 0,
			result: [][]int{
				{-8, 2, 6},
				{-8, 3, 5},
				{-6, 1, 5},
			},
		},
		"[3, 2, 1] target 6 should return triplet": {
			data:   []int{3, 2, 1},
			target: 6,
			result: [][]int{
				{1, 2, 3},
			},
		},
		"[1, 2, 3] target 7 should return triplet": {
			data:   []int{1, 2, 3},
			target: 7,
			result: [][]int{},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.ThreeNumSum(test.data, test.target)
			require.Equal(t, test.result, res)
		})
	}
}
