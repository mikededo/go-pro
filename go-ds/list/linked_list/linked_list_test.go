package linkedlist

import (
	"go-ds/helpers"
	"testing"
)

func initBaseList() *LL[int] {
	l := NewLinkedList[int]()
	l.Push(0).Push(1).Push(2)
	return l
}

func wantedAtPosition(t *testing.T, pos, got, want int) {
	t.Errorf("got %d @ pos %d, wanted %d", got, pos, want)
}

func validatePushedValues(t *testing.T, l *LL[int]) {
	if v, _ := l.At(0); v != 0 {
		wantedAtPosition(t, 0, v, 0)
	}
	if v, _ := l.At(1); v != 1 {
		wantedAtPosition(t, 1, v, 1)
	}
	if v, _ := l.At(2); v != 2 {
		wantedAtPosition(t, 2, v, 2)
	}
}

func validatePoppedValue(t *testing.T, l *LL[int], pos, want int) {
	v, _ := l.At(l.Length - 1)
	if v != want {
		wantedAtPosition(t, pos, v, want)
	}
}

func TestAt(t *testing.T) {
	Back := &LLNode[int]{Value: 1, Next: nil}
	Front := &LLNode[int]{Value: 0, Next: Back}
	l := &LL[int]{
		Front:  Front,
		Back:   Back,
		Length: 2,
	}

	t.Run("should return the element at the wanted position", func(t *testing.T) {
		v, err := l.At(1)
		if v != 1 {
			wantedAtPosition(t, 1, v, 1)
		}
		if err != nil {
			t.Error("got error, wanted nil")
		}
	})

	t.Run("should return error for accessing out of bounds", func(t *testing.T) {
		v, err := l.At(2)
		if v != *new(int) {
			t.Errorf("should have returned error, got %d", v)
		}
		if err == nil {
			t.Error("wanted error, got nil")
		}
	})
}

func TestPush(t *testing.T) {
	t.Run("should perform unchained push", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.Push(0)
		l.Push(1)
		l.Push(2)

		helpers.ValidateLength(t, 3, l.Length)
		validatePushedValues(t, l)
	})

	t.Run("should perform chained push", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.Push(0).Push(1).Push(2)

		helpers.ValidateLength(t, 3, l.Length)
		validatePushedValues(t, l)
	})
}

func TestPop(t *testing.T) {
	t.Run("should perform unchained pop", func(t *testing.T) {
		l := initBaseList()

		l.Pop()
		validatePoppedValue(t, l, 1, 1)
		helpers.ValidateLength(t, 2, l.Length)

		l.Pop()
		validatePoppedValue(t, l, 0, 0)
		helpers.ValidateLength(t, 1, l.Length)
	})

	t.Run("should perform chained pop", func(t *testing.T) {
		l := initBaseList()

		l.Pop().Pop()
		validatePoppedValue(t, l, 0, 0)
		helpers.ValidateLength(t, 1, l.Length)
	})
}

func TestDelete(t *testing.T) {
	t.Run("should delete @ first position", func(t *testing.T) {
		l := initBaseList()

		l.Delete(0)
		helpers.ValidateLength(t, 2, l.Length)
		v, _ := l.At(0)
		if v != 1 {
			wantedAtPosition(t, 0, v, 1)
		}
	})

	t.Run("should delete @ last position", func(t *testing.T) {
		l := initBaseList()

		l.Delete(2)
		_, err := l.At(2)
		if err == nil {
			t.Error("wanted error, got nil")
		}
	})

	t.Run("should delete @ any middle position", func(t *testing.T) {
		l := initBaseList()

		l.Delete(1)
		helpers.ValidateLength(t, 2, l.Length)
		v, _ := l.At(1)
		if v != 2 {
			wantedAtPosition(t, 1, v, 2)
		}
	})
}

func TestEach(t *testing.T) {
	t.Run("should apply the function to each element", func(t *testing.T) {
		l := initBaseList()

		l.Each(func(n *LLNode[int], _ int) {
			n.Value = n.Value * 2
		})
		l.Each(func(n *LLNode[int], i int) {
			if i*2 != n.Value {
				wantedAtPosition(t, i, n.Value, i*2)
			}
		})
	})
}

func TestIntegration(t *testing.T) {
	t.Run("should perform operations with the linked list", func(t *testing.T) {
		l := NewLinkedList[int]()
		l.Push(0).Push(1).Push(2).Push(3)

		if !l.Delete(2) {
			t.Error("expected to delete value 2 but did not")
		}
		if l.Delete(2) {
			t.Error("expected no to delete value 2 but did")
		}

		l.Pop()
		helpers.ValidateLength(t, 2, l.Length)
		l.Push(4)

		v, _ := l.At(l.Length - 1)
		if v != 4 {
			wantedAtPosition(t, 3, v, 4)
		}
	})
}
