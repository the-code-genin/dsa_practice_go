package data_structures

// A binary search tree
type BinarySearchTree struct {
	root *BinaryTreeNode
}

// Insert an element into the tree
func (bt *BinarySearchTree) Insert(key []byte, data []byte) error {
	node := NewBinaryTreeNode(key, data)

	// The root is not an actual node in the tree, just a placeholder
	if bt.root.right == nil {
		node.parent = bt.root
		bt.root.right = node
		return nil
	} else {
		return bt.root.right.Add(node)
	}
}

// Get an element's data from the tree with the element's key
func (bt *BinarySearchTree) Get(key []byte) ([]byte, error) {
	// Root's right child doesn't exist
	if bt.root.right == nil {
		return nil, BinaryTreeKeyNotFoundError{}
	}

	return bt.root.right.Get(key)
}

// Update an element's data from the tree with the element's key
func (bt *BinarySearchTree) Update(key, data []byte) error {
	// Root's right child doesn't exist
	if bt.root.right == nil {
		return BinaryTreeKeyNotFoundError{}
	}

	return bt.root.right.Update(key, data)
}

// Delete an element's data from the tree with the element's key
func (bt *BinarySearchTree) Delete(key []byte) error {
	// Root's right child doesn't exist
	if bt.root.right == nil {
		return BinaryTreeKeyNotFoundError{}
	}

	return bt.root.right.Delete(key)
}

// Create a new binary search tree
func NewBinarySearchTree() *BinarySearchTree {
	return &BinarySearchTree{NewBinaryTreeNode(nil, nil)}
}
