package algorithms

type SelectionSort struct{}

func (s SelectionSort) SortInt(input []int) []int {
	output := input

	for i := 0; i < len(output); i++ {
		smallestValueIndex, smallestValue := i, output[i]

		for j := i + 1; j < len(output); j++ {
			if output[j] < smallestValue {
				smallestValueIndex, smallestValue = j, output[j]
			}
		}

		if smallestValueIndex != i {
			output[i], output[smallestValueIndex] = smallestValue, output[i]
		}
	}

	return output
}
