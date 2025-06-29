package binary_tree

import (
	"dsa/queue"
	"fmt"
)

type Node struct {
	left *Node
	right *Node
	parent *Node
	value int
}

func NewNode(value int) *Node {
	return &Node{
		value: value,
	}
}

func (n *Node) TraversePreorder() {
	if n == nil {
		return
	}

	fmt.Println(n.value)
	n.left.TraversePreorder()
	n.right.TraversePreorder()
}

func (n *Node) TraverseInorder() {
	if n == nil {
		return
	}

	n.left.TraverseInorder()
	fmt.Println(n.value)
	n.right.TraverseInorder()
}

func (n *Node) TraversePostorder() {
	if n == nil {
		return
	}

	n.left.TraversePostorder()
	n.right.TraversePostorder()
	fmt.Println(n.value)
}

func (n *Node) TraverseBFS() {
	q := &queue.SliceQueue[*Node]{}
	q.Push(n)
	for item, err := q.Pop(); err == nil; item, err = q.Pop() {
		fmt.Println(item.value)
		if item.left != nil {
			q.Push(item.left)
		}
		if item.right != nil {
			q.Push(item.right)
		}
	}
}
