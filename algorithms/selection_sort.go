package algorithms

type SelectionSort struct {}

func (s SelectionSort) SortInt(input []int) ([]int) {
	output := input

	for i := 0; i < len(output); i++ {
		smallestValue := output[i]
		smallestIndex := i

		for j := i; j < len(output); j++ {
			if output[j] < smallestValue {
				smallestValue = output[j]
				smallestIndex = j
			}
		}

		if smallestIndex != i {
			oldSmallestValue := output[i]
	
			output[i] = smallestValue
			output[smallestIndex] = oldSmallestValue
		}
	}

	return output
}
