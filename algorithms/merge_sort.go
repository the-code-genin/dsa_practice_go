package algorithms

import (
	"math"
)

type MergeSort struct{}

func (s MergeSort) SortInt(partA, partB []int) []int {
	output := make([]int, 0)

	// Atomic sub array
	if len(partA) == 0 && len(partB) == 0 {
		return output
	} else if len(partA) == 1 && len(partB) == 0 {
		return partA
	} else if len(partB) == 1 && len(partA) == 0 {
		return partB
	}

	// Sort part a
	midPointA := int(math.Ceil(float64(len(partA)) / 2))
	sortedA := s.SortInt(partA[:midPointA], partA[midPointA:])

	// Sort part b
	midPointB := int(math.Ceil(float64(len(partB)) / 2))
	sortedB := s.SortInt(partB[:midPointB], partB[midPointB:])

	// Merge arrays
	for i, j := 0, 0; i < len(sortedA) || j < len(sortedB); {
		currA, currB := math.MaxInt, math.MaxInt
		if i < len(sortedA) {
			currA = sortedA[i]
		}
		if j < len(sortedB) {
			currB = sortedB[j]
		}

		if currA <= currB {
			output = append(output, currA)
			i++
		} else {
			output = append(output, currB)
			j++
		}
	}

	return output
}
