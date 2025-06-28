package main

import (
	"errors"
	"fmt"
)

func main() {
	var sq SliceQueue
	err := sq.Delete()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sq.Search(1))
	sq.Insert(5)
	fmt.Println(sq)
	sq.Insert(6)
	fmt.Println(sq)
	fmt.Println(sq.Search(5))
	fmt.Println(sq.Search(6))
	fmt.Println(sq.Search(7))
	err = sq.Delete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sq)
	err = sq.Delete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sq)
	err = sq.Delete()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sq)
}

type SliceQueue struct {
	items []int
}

var _ Queue = &SliceQueue{}

var ErrQueueEmpty = errors.New("queue is empty")

func (q *SliceQueue) Insert(value int) {
	q.items = append(q.items, value)
}

func (q *SliceQueue) Delete() error {
	if len(q.items) == 0 {
		return ErrQueueEmpty
	}

	q.items = q.items[1:]

	return nil
}

func (q *SliceQueue) Search(value int) bool {
	for _, v := range q.items {
		if v == value {
			return true
		}
	}

	return false
}

func (q *SliceQueue) String() string {
	return fmt.Sprint(q.items)
}

type Queue interface {
	Search(int) bool
	Insert(int)
	Delete() error
}
