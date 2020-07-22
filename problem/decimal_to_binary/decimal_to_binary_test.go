package decimaltobinary_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	dtb "code-interview/problem/decimal_to_binary"
)

func TestDecimalToBinary(t *testing.T) {
	tests := map[string]struct {
		num    int
		result string
	}{
		"3 should convert to 11": {
			num:    3,
			result: "11",
		},
		"5 should convert to 101": {
			num:    5,
			result: "101",
		},
		"9 should convert to 1001": {
			num:    9,
			result: "1001",
		},
	}

	for testName, test := range tests {
		t.Run(testName, func(t *testing.T) {
			res := dtb.DecimalToBinary(test.num)
			require.Equal(t, test.result, res)
		})
	}
}
