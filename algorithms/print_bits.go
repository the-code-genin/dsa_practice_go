package algorithms

import "math"

func PrintBits(input byte) []int8 {
	output := make([]int8, 0)

	for i := float64(7); i >= 0; i-- {
		if int8(input) & int8(math.Pow(2, i)) != 0 { // Bit is present
			output = append(output, 1)
		} else { // Bit is absent
			output = append(output, 0)
		}
	}

	return output
}