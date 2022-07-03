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

// Binary tree node doesn't have a parent error type
type BinaryTreeKeyParentlessError struct{}

func (err BinaryTreeKeyParentlessError) Error() string {
	return "binary tree node does not have a parent"
}

// Binary tree node doesn't have a successor error type
type BinaryTreeKeyNoSuccessorError struct{}

func (err BinaryTreeKeyNoSuccessorError) Error() string {
	return "binary tree node does not have a successor"
}

// A binary tree node
type BinaryTreeNode struct {
	key    []byte
	data   []byte
	parent *BinaryTreeNode
	left   *BinaryTreeNode
	right  *BinaryTreeNode
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
			btn.right.parent = btn
			return nil
		}

		// Tell the right node to add the new node
		return btn.right.Add(node)
	} else { // New node key is lesser than the current node key
		if btn.left == nil { // If the left leaf is empty, set it to the new node
			btn.left = node
			btn.left.parent = btn
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
			if btn.right != nil { // If the right leaf is not empty
				// Tell the right node to get the data based on the search key
				return btn.right.Get(key)
			}
		} else { // The search key is lesser than the current node key
			if btn.left != nil { // If the left leaf is not empty
				// Tell the left node to get the data based on the search key
				return btn.left.Get(key)
			}
		}

		return nil, BinaryTreeKeyNotFoundError{}
	}
}

// Update data by comparing a node's key with the search key
// If the node's key doesn't match, it propagates the key search to the appropriate child
// If the node doesn't have appropriate children, it returns an error
func (btn *BinaryTreeNode) Update(key []byte, data []byte) error {
	tnKey := new(big.Int).SetBytes(btn.key)
	nodeKey := new(big.Int).SetBytes(key)

	if x := tnKey.Cmp(nodeKey); x == 0 { // The right key
		btn.data = data
		return nil
	} else { // Not the right key
		if x > 0 { // The search key is greater than current node key
			if btn.right != nil { // If the right leaf is not empty
				// Tell the right node to update the data based on the search key
				return btn.right.Update(key, data)
			}
		} else { // The search key is lesser than the current node key
			if btn.left != nil { // If the left leaf is not empty
				// Tell the left node to get the data based on the search key
				return btn.left.Update(key, data)
			}
		}

		return BinaryTreeKeyNotFoundError{}
	}
}

// Find the inorder successor to a key
func (btn *BinaryTreeNode) FindInorderSuccessor(key []byte) (*BinaryTreeNode, error) {
	tnKey := new(big.Int).SetBytes(btn.key)
	nodeKey := new(big.Int).SetBytes(key)

	// This node is not a leaf node, propagate to it's children
	if btn.left != nil || btn.right != nil {
		if x := tnKey.Cmp(nodeKey); x > 0 { // Propagate to it's left child if possible since it has a greater key
			if btn.left != nil {
				return btn.left.FindInorderSuccessor(key)
			}
		} else { // Propagate to it's right children if possible since it has a lesser key
			if btn.right != nil {
				return btn.right.FindInorderSuccessor(key)
			}
		}

		return nil, BinaryTreeKeyNoSuccessorError{}
	}

	// Since it's a leaf node return it
	return btn, nil
}

// Delete data by comparing a node's key with the search key
// If the node's key doesn't match, it propagates the key search to the appropriate child
// If the node doesn't have appropriate children, it returns an error
// On deletion;
// If the node is a leaf node, delete the node from it's parent
// If the node has only one child, that child is used to replace the node in the tree
// If the node has two children, find the successor node from the children's leaves
func (btn *BinaryTreeNode) Delete(key []byte) error {
	tnKey := new(big.Int).SetBytes(btn.key)
	nodeKey := new(big.Int).SetBytes(key)

	// Root node with no parent from it's implementation data structure i.e the Set, Binary Search Tree or Hash Map
	if btn.parent == nil {
		return BinaryTreeKeyParentlessError{}
	}

	if x := tnKey.Cmp(nodeKey); x == 0 { // The right key
		if btn.left == nil && btn.right == nil { // Node is a leaf node, remove from parent without any issues
			if btn.parent.left == btn {
				btn.parent.left = nil
			} else {
				btn.parent.right = nil
			}
		} else if (btn.left != nil && btn.right == nil) || (btn.right != nil && btn.left == nil) { // Node has only one child node
			if btn.left != nil && btn.right == nil { // Left node is not empty
				if btn.parent.left == btn {
					btn.parent.left = btn.left
				} else {
					btn.parent.right = btn.left
				}
			} else { // Right node is not empty
				if btn.parent.left == btn {
					btn.parent.left = btn.right
				} else {
					btn.parent.right = btn.right
				}
			}
		} else { // Node has two children
			// Find the successor node from both branches of the node
			successor, err := btn.right.FindInorderSuccessor(btn.key)
			if err != nil {
				successor, err = btn.left.FindInorderSuccessor(btn.key)
				if err != nil {
					return err
				}
			}

			// Tell successor to delete itself from it's parent
			err = successor.Delete(successor.key)
			if err != nil {
				return err
			}

			// Update the successor's meta data to match the current node
			successor.left, successor.right = btn.left, btn.right
			if btn.parent.left == btn {
				btn.parent.left = successor
			} else {
				btn.parent.right = successor
			}
		}

		return nil
	} else { // Not the right key
		if x > 0 { // The search key is greater than current node key
			if btn.right != nil { // If the right leaf is not empty
				// Tell the right node to get the data based on the search key
				return btn.right.Delete(key)
			}
		} else { // The search key is lesser than the current node key
			if btn.left != nil { // If the left leaf is not empty
				// Tell the left node to get the data based on the search key
				return btn.left.Delete(key)
			}
		}

		return BinaryTreeKeyNotFoundError{}
	}
}

// Initialize a new binary tree node without any links
func NewBinaryTreeNode(key, data []byte) *BinaryTreeNode {
	return &BinaryTreeNode{key, data, nil, nil, nil}
}
