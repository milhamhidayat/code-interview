package program_test

import (
	"testing"

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
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := p.KadaneAlgorithm(test.data)
			require.Equal(t, test.result, res)
		})
	}
}
