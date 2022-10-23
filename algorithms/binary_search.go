package algorithms

import "math"

func binarySearch(elems []int, elem, low, high int) int {
	newHigh := high
	newLow := low
	midPoint := (float64(newHigh-newLow) / 2) + float64(newLow)

	if float64(int(midPoint)) != midPoint { // Even no of elements
		midPoint1 := int(math.Floor(midPoint))
		midPoint2 := int(math.Ceil(midPoint))

		midElem1 := elems[midPoint1]
		midElem2 := elems[midPoint2]

		if midElem1 == elem {
			return midPoint1
		} else if midElem2 == elem {
			return midPoint2
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
				newLow = int(midPoint)
			} else {
				newHigh = int(midPoint)
			}
		}
	}

	if newHigh == newLow {
		return -1
	}

	return binarySearch(elems, elem, newLow, newHigh)
}

func BinarySearch(elems []int, elem int) int {
	if len(elems) == 0 { // Empty array
		return -1
	}

	return binarySearch(elems, elem, 0, len(elems)-1)
}
