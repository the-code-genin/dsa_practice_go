package data_structures

import (
	"testing"
)

func TestBinarySearchTreeInitialization(t *testing.T) {
	cases := []struct {
		key  string
		data string
	}{
		{ "hello", "world" },
		{ "foo", "bar" },
		{ "mohammed", "adekunle" },
		{ "lorem", "ipsum" },
		{ "code", "genin" },
		{ "genghis", "khan" },
	}

	// Initialize the tree
	tree := NewBinarySearchTree()
	for _, tc := range cases {
		if err := tree.Insert([]byte(tc.key), []byte(tc.data)); err != nil {
			t.Error(err)
		}
	}

	// Confirm the key data matches while retrieving
	for _, tc := range cases {
		res, err := tree.Get([]byte(tc.key))
		if err != nil {
			t.Error(err)
		} else if string(res) != tc.data {
			t.Errorf("expected %v, got %v", tc.data, string(res))
		}
	}
}