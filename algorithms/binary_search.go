package algorithms

import "math"

type BinarySearch struct{}

func (s BinarySearch) Search(elems []int, elem int, low, high uint) int {
	newHigh := high
	newLow := low
	midPoint := (float64(newHigh-newLow) / 2) + float64(newLow)

	if float64(int(midPoint)) != midPoint { // Even no of elements
		midPoint1 := uint(math.Floor(midPoint))
		midPoint2 := uint(math.Ceil(midPoint))

		midElem1 := elems[midPoint1]
		midElem2 := elems[midPoint2]

		if midElem1 == elem {
			return int(midPoint1)
		} else if midElem2 == elem {
			return int(midPoint2)
		} else {
			if elem > midElem2 {
				newLow = midPoint2
			} else {
				newHigh = midPoint1
			}
		}
	} else { // Odd no of elements
		if midElem := elems[int(midPoint)]; midElem == elem {
			return int(midPoint)
		} else {
			if elem > midElem {
				newLow = uint(midPoint)
			} else {
				newHigh = uint(midPoint)
			}
		}
	}

	if newHigh == newLow {
		return -1
	}

	return s.Search(elems, elem, newLow, newHigh)
}
