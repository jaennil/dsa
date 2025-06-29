package main

import (
	"errors"
	"fmt"
)

func main() {
	ss := &SliceStack{}
	fmt.Println(ss)
	err := ss.Pop()
	if err != nil {
		fmt.Println(err)
	}
	ss.Push(2)
	ss.Push(3)
	ss.Push(4)
	fmt.Println(ss)

	err = ss.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ss)

	err = ss.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ss)

	err = ss.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ss)

	err = ss.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(ss)
}

type SliceStack struct {
	items []int
}

var ErrStackEmpty = errors.New("stack is empty")

func (s *SliceStack) Push(value int) {
	s.items = append(s.items, value)
}

func (s *SliceStack) Pop() error {
	if len(s.items) == 0 {
		return ErrStackEmpty
	}

	s.items = s.items[:len(s.items)-1]

	return nil
}

var _ Stack = &SliceStack{}

type Stack interface {
	Push(int)
	Pop() error
}
