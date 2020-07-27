package program_test

import (
	"testing"

	p "code-interview/algoexpert/easy/find_three_largest_numbers"

	"github.com/stretchr/testify/require"
)

func TestFindThreeLargestNumbers(t *testing.T) {
	tests := map[string]struct {
		data   []int
		result []int
	}{
		"[5, 1, 7, 3, 4] should return [4, 5, 7]": {
			data:   []int{5, 1, 7, 3, 4},
			result: []int{4, 5, 7},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.FindThreeLargestNumbers(test.data)
			require.Equal(t, test.result, res)
		})
	}

}
