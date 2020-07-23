package program_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	p "code-interview/algoexpert/easy/binary_search"
)

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		data   []int
		num    int
		result int
	}{
		"[1, 5, 23, 111] num 111 should return 3": {
			data:   []int{1, 5, 23, 111},
			num:    111,
			result: 3,
		},
		"[1, 5, 23, 111] num 120 should return -1": {
			data:   []int{1, 5, 23, 111},
			num:    120,
			result: -1,
		},
		"[0, 1, 21, 33, 45, 45, 61, 71, 72, 73] num 72 should return 8": {
			data:   []int{0, 1, 21, 33, 45, 45, 61, 71, 72, 73},
			num:    72,
			result: 8,
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.BinarySearch(test.data, test.num)
			require.Equal(t, test.result, res)
		})
	}
}
