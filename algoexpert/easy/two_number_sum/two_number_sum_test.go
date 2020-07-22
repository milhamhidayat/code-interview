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
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.TwoNumberSum(test.data, test.result)
			require.Equal(t, test.result, res)
		})
	}
}
