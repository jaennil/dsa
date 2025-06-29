package binary_tree

import (
	"fmt"
	"testing"
)

func TestBinaryTree(t *testing.T) {
	root := NewNode(1)
	root.left = NewNode(2)
	root.right = NewNode(3)
	root.left.left = NewNode(4)
	root.left.right = NewNode(5)
	root.right.right = NewNode(6)
	root.TraversePreorder()
	fmt.Println()
	root.TraverseInorder()
	fmt.Println()
	root.TraversePostorder()
	fmt.Println()
	root.TraverseBFS()
}
