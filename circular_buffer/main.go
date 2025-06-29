package main

import (
	"errors"
	"fmt"
)

func main() {
	scb := NewSliceCircularBuffer(3)
	scb.Push(1)
	fmt.Println(scb)
	scb.Push(2)
	fmt.Println(scb)
	scb.Push(3)
	fmt.Println(scb)
	scb.Push(4)
	fmt.Println(scb)
	scb.Push(5)
	fmt.Println(scb)
	var err error
	err = scb.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(scb)
	err = scb.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(scb)
	err = scb.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(scb)
	err = scb.Pop()
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(scb)
}

type SliceCircularBuffer struct {
	readIndex int
	writeIndex int
	len int
	items []int
}

var ErrBufferEmpty = errors.New("buffer is empty")

func NewSliceCircularBuffer(size int) *SliceCircularBuffer {
	return &SliceCircularBuffer{
		readIndex: 0,
		writeIndex: 0,
		len: 0,
		items: make([]int, size),
	}
}

func (b *SliceCircularBuffer) Push(value int) {
	if b.writeIndex == cap(b.items) {
		b.writeIndex = 0
	}

	b.items[b.writeIndex] = value
	b.writeIndex++

	if b.len < cap(b.items) {
		b.len++
	}
}

func (b *SliceCircularBuffer) Pop() error {
	if b.len == 0 {
		return ErrBufferEmpty
	}

	b.len--
	b.readIndex++

	return nil
}
