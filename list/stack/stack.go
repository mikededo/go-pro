package stack

import "errors"

type Stack[K comparable] struct {
	bottom *SNode[K]
	top    *SNode[K]
	length int
}

type SNode[K comparable] struct {
	Next  *SNode[K]
	Value K
}

func NewStack[K comparable]() *Stack[K] {
	return &Stack[K]{}
}

func newNode[K comparable](v K) *SNode[K] {
	return &SNode[K]{Value: v}
}

func (s *Stack[K]) Push(v K) *Stack[K] {
	n := newNode(v)

	if s.bottom == nil {
		s.bottom = n
	} else {
		s.top.Next = n
	}
	s.top = n
	s.length++

	return s
}

func (s *Stack[K]) Pop() *Stack[K] {
	if !s.Empty() {
		if s.top == s.bottom {
			s.top, s.bottom = nil, nil
		} else {
			c := s.bottom
			for c != s.top {
				c = c.Next
			}
			c.Next = nil
		}
		s.length--

		return s
	}

	return s
}

func (s *Stack[K]) Empty() bool {
	return s.bottom == nil
}

func (s *Stack[K]) Top() (K, error) {
	if s.Empty() {
		return *new(K), errors.New("stack is empty")
	}

	return s.top.Value, nil
}

func (s *Stack[K]) Length() int {
	return s.length
}
