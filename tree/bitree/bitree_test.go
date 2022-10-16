package bitree

import (
	"testing"
)

func generateBaseTree() *BTree[int] {
	b := NewBTree(less)

	b.Insert(5)
	b.Insert(3)
	b.Insert(1)
	b.Insert(7)
	b.Insert(6)
	b.Insert(8)

	return b
}

func less(a, b int) bool { return a < b }

func validateInorder(t *testing.T, wanted, got string) {
	if got != wanted {
		t.Errorf("wanted %s for inorder, but got %s", wanted, got)
	}
}

func TestInorder(t *testing.T) {
	t.Run("should return empty string on traversing empty tree", func(t *testing.T) {
		b := NewBTree(less)
		r := b.Inorder()
		if r != "" {
			t.Errorf("expected inorder to return empty string, but got %s", r)
		}
	})

	t.Run("should traverse the tree in inorder", func(t *testing.T) {
		b := generateBaseTree()
		validateInorder(t, "135678", b.Inorder())
	})
}

func TestInsert(t *testing.T) {
	t.Run("should properly insert the values", func(t *testing.T) {
		b := generateBaseTree()

		validateInorder(t, "135678", b.Inorder())
	})

	t.Run("should always insert to the left", func(t *testing.T) {
		b := NewBTree(less)

		b.Insert(5)
		b.Insert(4)
		b.Insert(3)
		b.Insert(2)
		b.Insert(1)

		if b.root.R != nil {
			t.Errorf("right children should be empty, but got %d", b.root.R.Value)
		}
		validateInorder(t, "12345", b.Inorder())
	})

	t.Run("should always insert to the right", func(t *testing.T) {
		b := NewBTree(less)

		b.Insert(1)
		b.Insert(2)
		b.Insert(3)
		b.Insert(4)
		b.Insert(5)

		if b.root.L != nil {
			t.Errorf("left children should be empty, but got %d", b.root.L.Value)
		}
		validateInorder(t, "12345", b.Inorder())
	})
}

func TestDelete(t *testing.T) {
	t.Run("should properly delete a leaf", func(t *testing.T) {
		b := generateBaseTree()

		b.Delete(8)
		validateInorder(t, "13567", b.Inorder())
	})

	t.Run("should properly delete the root", func(t *testing.T) {
		b := generateBaseTree()

		b.Delete(5)
		validateInorder(t, "13678", b.Inorder())
	})

	t.Run("should properly delete a node with one child", func(t *testing.T) {
		b := generateBaseTree()

		b.Delete(1)
		validateInorder(t, "35678", b.Inorder())
	})

	t.Run("should not do anything if tree is empty", func(t *testing.T) {
		b := NewBTree(less)
		b.Delete(8)

		if b.root != nil {
			t.Errorf("expected root to be nil, got %v", b.root)
		}
	})
}

func TestMin(t *testing.T) {
	t.Run("should return the minimum value of the tree", func(t *testing.T) {
		m, e := generateBaseTree().Min()
		if m != 1 {
			t.Errorf("wanted %d as min value, got %d", 1, m)
		}
		if e != nil {
			t.Errorf("got %s as error, but wanted nil", e)
		}
	})

	t.Run("should return error if the tree is empty", func(t *testing.T) {
		_, e := NewBTree(less).Min()
		if e == nil {
			t.Errorf("expected \"%s\", but got nil", "tree is empty")
		}
	})
}

func TestMax(t *testing.T) {
	t.Run("should return the maximum value of the tree", func(t *testing.T) {
		m, e := generateBaseTree().Max()
		if m != 8 {
			t.Errorf("wanted %d as max value, got %d", 8, m)
		}
		if e != nil {
			t.Errorf("got %s as error, but wanted nil", e)
		}
	})

	t.Run("should return error if the tree is empty", func(t *testing.T) {
		_, e := NewBTree(less).Max()
		if e == nil {
			t.Errorf("expected \"%s\", but got nil", "tree is empty")
		}
	})
}

func TestContains(t *testing.T) {
	t.Run("should return false if tree is empty", func(t *testing.T) {
		if (NewBTree(less)).Contains(1) {
			t.Error("expected empty tree not to contain 1")
		}
	})

	t.Run("should return false if tree does not contain the given value", func(t *testing.T) {
		if generateBaseTree().Contains(9) {
			t.Error("expected tree not to contain 9")
		}
	})

	t.Run("should return true if tree does contain a value higher than the root", func(t *testing.T) {
		if !generateBaseTree().Contains(8) {
			t.Error("expected tree to contain 8")
		}
	})

	t.Run("should return true if tree does contain a value lower than the root", func(t *testing.T) {
		if !generateBaseTree().Contains(1) {
			t.Error("expected tree to contain 1")
		}
	})
}
