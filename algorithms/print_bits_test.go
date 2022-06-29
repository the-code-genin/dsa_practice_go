package algorithms

import "testing"

func compareSlice(a, b []int8) bool {
	if len(a) != len(b) {
		return false
	}

	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}

	return true
}

func TestPrintBits(t *testing.T) {
	cases := []struct {
		in  byte
		out []int8
	}{
		{1, []int8{0, 0, 0, 0, 0, 0, 0, 1}},
		{3, []int8{0, 0, 0, 0, 0, 0, 1, 1}},
		{0, []int8{0, 0, 0, 0, 0, 0, 0, 0}},
		{4, []int8{0, 0, 0, 0, 0, 1, 0, 0}},
		{2, []int8{0, 0, 0, 0, 0, 0, 1, 0}},
		{5, []int8{0, 0, 0, 0, 0, 1, 0, 1}},
		{255, []int8{1, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, tc := range cases {
		res := PrintBits(tc.in)
		if !compareSlice(res, tc.out) {
			t.Errorf("expected %v, got %v", tc.out, res)
		}
	}
}
