package linkedlist

import (
	"errors"
	"fmt"
	"strings"
)

type LL[K comparable] struct {
	Length int
	Front  *LLNode[K]
	Back   *LLNode[K]
}

type LLNode[K comparable] struct {
	Next  *LLNode[K]
	Value K
}

func NewLinkedList[K comparable]() *LL[K] {
	return &LL[K]{
		Front:  nil,
		Back:   nil,
		Length: 0,
	}
}

// Returns the element at the given index
func (l *LL[K]) At(i int) (K, error) {
	if i >= l.Length {
		return *new(K), errors.New(strings.Join([]string{"index out of bounds [0:", fmt.Sprint(l.Length), "]"}, ""))
	}

	c := l.Front
	j := 0
	for j < i {
		c = c.Next
		j++
	}

	return c.Value, nil
}

// Adds an item to the given linked list and returns the same
// linked list with the added items
func (l *LL[K]) Push(v K) *LL[K] {
	next := &LLNode[K]{Next: nil, Value: v}

	if l.Front == nil {
		l.Front = next
		l.Back = next
	} else if l.Front != nil {
		l.Back.Next = next
		l.Back = next
	}

	l.Length++
	return l
}

// Removes the last item and returns the same linked list
// with the last item removed
func (l *LL[K]) Pop() *LL[K] {
	if l.Length == 0 {
		return l
	}

	c := l.Front
	for c.Next != l.Back {
		c = c.Next
	}
	l.Back = c
	l.Length--

	return l
}

// Removes the item if exists, returning a bool on whether the
// items have been removed or not
func (l *LL[K]) Delete(v K) bool {
	if l.Length == 0 {
		return false
	}

	if l.Front.Value == v {
		l.Front = l.Front.Next
		l.Length--
		return true
	}

	c := l.Front.Next
	p := l.Front
	for c.Next != nil && c.Value != v {
		p = c
		c = c.Next
	}

	if c.Next != nil {
		p.Next = c.Next
	} else if c.Next == nil && c.Value == v {
		l.Back = p
	} else {
		return false
	}

	l.Length--
	return true
}

// Runs the given function to each item of the list
func (l *LL[K]) Each(fn func(*LLNode[K], int)) {
	c := l.Front
	i := 0

	for c != nil {
		fn(c, i)

		c = c.Next
		i++
	}
}
