package algorithms

import "testing"

func TestBinarySearch(t *testing.T) {
	cases := []struct{
		elems []int
		elem int
		index int
	}{
		{ []int{1, 2, 3, 4, 5, 6, 7}, 4, 3 },
		{ []int{3, 6, 9, 10}, 6, 1 },
		{ []int{9, 10}, 9, 0 },
		{ []int{6, 9, 12, 13, 16}, 11, -1 },
	}

	for _, tc := range cases {
		res := BinarySearch(tc.elems, tc.elem)
		if res != tc.index {
			t.Errorf("expected %v, got %v", tc.index, res)
		}
	}
}