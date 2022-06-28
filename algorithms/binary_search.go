package algorithms

func binarySearch(elems []int, elem, low, high int) int {
	newHigh := high
	newLow := low

	if (high + 1 - low) % 2 != 0 { // Odd no of elements
		midPoint := ((high - 1) / 2) + low
		midElem := elems[midPoint]

		if midElem == elem {
			return midPoint
		} else {
			if elem > midElem {
				newLow = midPoint
			} else {
				newHigh = midPoint
			}
		}
	} else { // Even no of elements
		midPoint1 := ((high - 1) / 2) + low
		midPoint2 := midPoint1 + 1

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

	return binarySearch(elems, elem, 0, len(elems) - 1)
}