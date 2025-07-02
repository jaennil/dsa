package min_heap

import (
	"golang.org/x/exp/constraints"
)

type Node[T constraints.Ordered] struct {
	left *Node[T]
	right *Node[T]
	value T
}

func NewNode[T constraints.Ordered](value T) *Node[T] {
	return &Node[T]{value: value}
}

func (n *Node[T]) Insert(value T) {
	if value > n.value {

	}
}
