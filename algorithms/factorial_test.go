package algorithms

import "testing"

func TestFactorial(t *testing.T) {
	cases := []struct{
		in, out int64
	}{
		{ 2, 2 },
		{ 5, 120 },
		{ 8, 40320 },
	}

	for _, tc := range cases {
		res := Factorial(tc.in)
		if res != tc.out {
			t.Errorf("expected %v, got %v", tc.out, res)
		}
	}
}