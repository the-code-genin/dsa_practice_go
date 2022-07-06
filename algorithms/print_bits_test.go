package algorithms

import (
	"bytes"
	"testing"
)

func TestPrintBits(t *testing.T) {
	cases := []struct {
		in  byte
		out []byte
	}{
		{1, []byte{0, 0, 0, 0, 0, 0, 0, 1}},
		{3, []byte{0, 0, 0, 0, 0, 0, 1, 1}},
		{0, []byte{0, 0, 0, 0, 0, 0, 0, 0}},
		{4, []byte{0, 0, 0, 0, 0, 1, 0, 0}},
		{2, []byte{0, 0, 0, 0, 0, 0, 1, 0}},
		{5, []byte{0, 0, 0, 0, 0, 1, 0, 1}},
		{255, []byte{1, 1, 1, 1, 1, 1, 1, 1}},
	}

	for _, tc := range cases {
		res := PrintBits(tc.in)
		if !bytes.Equal(res, tc.out) {
			t.Errorf("expected %v, got %v", tc.out, res)
		}
	}
}
