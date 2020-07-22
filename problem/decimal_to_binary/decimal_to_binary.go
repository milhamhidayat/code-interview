package decimaltobinary

import "strconv"

// DecimalToBinary convert decimal to binary
func DecimalToBinary(num int) string {
	var result string
	for num > 0 {
		newNum := num / 2
		mod := num % 2
		result += strconv.Itoa(mod)
		num = newNum
	}

	return result
}
