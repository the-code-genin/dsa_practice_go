package data_structures

import (
	"math/big"
)

// Key collission error type
type BinaryTreeKeyCollissionError struct{}

func (err BinaryTreeKeyCollissionError) Error() string {
	return "key collission occured"
}

// Key not found error type
type BinaryTreeKeyNotFoundError struct{}

func (err BinaryTreeKeyNotFoundError) Error() string {
	return "key not found"
}

// A binary tree node
type BinaryTreeNode struct {
	key   []byte
	data  []byte
	left  *BinaryTreeNode
	right *BinaryTreeNode
}

// Add a new node to an existing tree node
func (btn *BinaryTreeNode) Add(node *BinaryTreeNode) error {
	tnKey := new(big.Int).SetBytes(btn.key)
	nodeKey := new(big.Int).SetBytes(node.key)

	if x := tnKey.Cmp(nodeKey); x == 0 { // Collission occurs
		return BinaryTreeKeyCollissionError{}
	} else if x > 0 { // New node key is greater than current node key
		if btn.right == nil { // If the right leaf is empty, set it to the new node
			btn.right = node
			return nil
		}

		// Tell the right node to add the new node
		return btn.right.Add(node)
	} else { // New node key is lesser than the current node key
		if btn.left == nil { // If the left leaf is empty, set it to the new node
			btn.left = node
			return nil
		}

		// Tell the left node to add the new node
		return btn.left.Add(node)
	}
}

// Get data by comparing a node's key with the search key
// If the node's key doesn't match, it propagates the key search to the appropriate child
// If the node doesn't have appropriate children, it returns an error
func (btn *BinaryTreeNode) Get(key []byte) ([]byte, error) {
	tnKey := new(big.Int).SetBytes(btn.key)
	nodeKey := new(big.Int).SetBytes(key)

	if x := tnKey.Cmp(nodeKey); x == 0 { // The right key
		return btn.data, nil
	} else { // Not the right key
		if x > 0 { // The search key is greater than current node key
			if btn.right == nil { // If the right leaf is empty
				return nil, BinaryTreeKeyNotFoundError{}
			}

			// Tell the right node to get the data based on the search key
			return btn.right.Get(key)
		} else { // The search key is lesser than the current node key
			if btn.left == nil { // If the left leaf is empty
				return nil, BinaryTreeKeyNotFoundError{}
			}

			// Tell the left node to get the data based on the search key
			return btn.left.Get(key)
		}
	}
}

func NewBinaryTreeNode(key, data []byte) *BinaryTreeNode {
	return &BinaryTreeNode{key, data, nil, nil}
}

// A binary search tree
type BinarySearchTree struct {
	root *BinaryTreeNode
}

// Insert an element into the tree
func (bt *BinarySearchTree) Insert(key []byte, data []byte) error {
	node := NewBinaryTreeNode(key, data)

	// Set the node as a root node if none exists or add to the root node
	// which can further propagate the node down the tree
	if bt.root == nil {
		bt.root = node
		return nil
	} else {
		return bt.root.Add(node)
	}
}

// Get an element's data from the tree with the element's key
func (bt *BinarySearchTree) Get(key []byte) ([]byte, error) {
	// Root doesn't exist
	if bt.root == nil {
		return nil, BinaryTreeKeyNotFoundError{}
	}

	return bt.root.Get(key)
}

func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{}
}
