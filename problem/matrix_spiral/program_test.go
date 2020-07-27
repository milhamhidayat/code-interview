package program_test

import (
	"testing"

	p "code-interview/problem/matrix_spiral"

	"github.com/stretchr/testify/require"
)

func TestMatrixSpiral(t *testing.T) {
	tests := map[string]struct {
		num int
		res [][]int
	}{
		"matrix 2": {
			num: 2,
			res: [][]int{
				{1, 2},
				{4, 3},
			},
		},
		"matrix 3": {
			num: 3,
			res: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.MatrixSpiral(test.num)
			require.Equal(t, test.res, res)
		})
	}
}
