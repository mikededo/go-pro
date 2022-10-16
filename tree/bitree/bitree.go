package bitree

import (
	"errors"
	"fmt"
)

type BNode[K comparable] struct {
	L     *BNode[K]
	R     *BNode[K]
	Value K
}

type BTree[K comparable] struct {
	root *BNode[K]
	// Less (<) comparator. Should return true if a is less than b
	Less func(a, b K) bool
}

func newBNode[K comparable](v K) *BNode[K] {
	return &BNode[K]{Value: v}
}

func NewBTree[K comparable](Less func(a, b K) bool) *BTree[K] {
	return &BTree[K]{Less: Less}
}

func (b *BTree[K]) Insert(v K) {
	if b.root == nil {
		b.root = newBNode(v)
	} else {
		b.insert(b.root, v)
	}
}

func (b *BTree[K]) insert(n *BNode[K], v K) *BNode[K] {
	if n == nil {
		return newBNode(v)
	} else if b.Less(v, n.Value) {
		n.L = b.insert(n.L, v)
	} else {
		n.R = b.insert(n.R, v)
	}

	return n
}

func (b *BTree[K]) Delete(v K) {
	if b.root == nil {
		return
	}

	b.root = b.delete(b.root, v)
}

func (b *BTree[K]) delete(n *BNode[K], v K) *BNode[K] {
	if n == nil {
		return nil
	}

	if n.Value == v {
		if n.L == nil && n.R == nil {
			return nil
		} else if n.R == nil {
			return n.L
		} else if n.L == nil {
			return n.R
		}

		// reorganize tree
		m := b.min(n.R)
		n.Value = m
		n.R = b.delete(n.R, m)
	} else if b.Less(v, n.Value) {
		n.L = b.delete(n.L, v)
	} else {
		// !b.Less(n.Value, v)
		n.R = b.delete(n.R, v)
	}

	return n
}

func (b *BTree[K]) Min() (K, error) {
	if b.root == nil {
		return *new(K), errors.New("tree is empty")
	}

	return b.min(b.root), nil
}

func (b *BTree[K]) min(n *BNode[K]) K {
	if n.L == nil {
		return n.Value
	}

	return b.min(n.L)
}

func (b *BTree[K]) Max() (K, error) {
	if b.root == nil {
		return *new(K), errors.New("tree is empty")
	}

	return b.max(b.root), nil
}

func (b *BTree[K]) max(n *BNode[K]) K {
	if n.R == nil {
		return n.Value
	}

	return b.max(n.R)
}

func (b *BTree[K]) Inorder() string {
	return b.inorder(b.root)
}

func (b *BTree[K]) inorder(n *BNode[K]) string {
	if n == nil {
		return ""
	}

	return b.inorder(n.L) + fmt.Sprint(n.Value) + b.inorder(n.R)
}

func (b *BTree[K]) Contains(v K) bool {
	if b.root == nil {
		return false
	}

	return b.contains(b.root, v)
}

func (b *BTree[K]) contains(n *BNode[K], v K) bool {
	if n == nil {
		return false
	} else if n.Value == v {
		return true
	} else if b.Less(v, n.Value) {
		return b.contains(n.L, v)
	}

	return b.contains(n.R, v)
}
