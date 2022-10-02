package algorithms

import (
	"testing"
)

func TestSeletionSort(t *testing.T) {
	testCases := []struct {
		in, out []int
	}{
		{[]int{2, 5, 1}, []int{1, 2, 5}},
		{[]int{9, 2, 5, 6, 1, 8}, []int{1, 2, 5, 6, 8, 9}},
	}

	for _, tc := range testCases {
		res := SelectionSort{}.SortInt(tc.in)
		if len(res) != len(tc.out) {
			t.Errorf("%v was expected to match %v", res, tc.out)
		}

		for i := range res {
			if res[i] != tc.out[i] {
				t.Errorf("%v was expected to match %v", res, tc.out)
			}
		}
	}
}
