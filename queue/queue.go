package queue

import (
	"errors"
)

type SliceQueue[T comparable] struct {
	items []T
}

func NewSliceQueue[T comparable]() *SliceQueue[T] {
	return &SliceQueue[T]{}
}

var _ Queue[int] = &SliceQueue[int]{}

var ErrQueueEmpty = errors.New("queue is empty")

func (q *SliceQueue[T]) Push(value T) {
	q.items = append(q.items, value)
}

func (q *SliceQueue[T]) Pop() (T, error) {
	if len(q.items) == 0 {
		// TODO: maybe pointer is better
		var zero T
		return zero, ErrQueueEmpty
	}

	item := q.items[0]

	q.items = q.items[1:]

	return item, nil
}

func (q *SliceQueue[T]) Search(value T) bool {
	for _, v := range q.items {
		if v == value {
			return true
		}
	}

	return false
}

func (q *SliceQueue[T]) IsEmpty() bool {
	return len(q.items) == 0
}

type Queue[T comparable] interface {
	Search(T) bool
	Push(T)
	Pop() (T, error)
	IsEmpty() bool
}
