package queue

import (
	"fmt"
	"testing"
)

func TestQueue(t *testing.T) {
	var sq SliceQueue[int]
	_, err := sq.Pop()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(sq.Search(1))
	sq.Push(5)
	fmt.Println(sq)
	sq.Push(6)
	fmt.Println(sq)
	fmt.Println(sq.Search(5))
	fmt.Println(sq.Search(6))
	fmt.Println(sq.Search(7))
	v, err := sq.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(v)
	fmt.Println(sq)
	_, err = sq.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sq)
	_, err = sq.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(sq)
}
