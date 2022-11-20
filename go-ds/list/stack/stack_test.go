package stack

import "testing"
import "go-ds/helpers"

func validateValue(t *testing.T, want, got int) {
	if got != want {
		t.Errorf("got %d, wanted %d", got, want)
	}
}

func TestLength(t *testing.T) {
	t.Run("should return 0 for empty stack", func(t *testing.T) {
		helpers.ValidateLength(t, 0, NewStack[int]().Length())
	})

	t.Run("should return size for non-empty stack", func(t *testing.T) {
		helpers.ValidateLength(t, 2, NewStack[int]().Push(1).Push(2).Length())
	})
}

func TestPush(t *testing.T) {
	t.Run("should perform unchained push", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1)
		s.Push(2)
		s.Push(3)

		helpers.ValidateLength(t, 3, s.Length())
		validateValue(t, 3, s.top.Value)
	})

	t.Run("should perform chained push", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1).Push(2).Push(3)

		helpers.ValidateLength(t, 3, s.Length())
		validateValue(t, 3, s.top.Value)
	})
}

func TestPop(t *testing.T) {
	t.Run("should not pop an empty stack", func(t *testing.T) {
		s := NewStack[int]()
		helpers.ValidateLength(t, 0, s.Length())

		s.Pop()
		helpers.ValidateLength(t, 0, s.Length())
	})

	t.Run("should perform unchained pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1).Push(2).Push(3)
		helpers.ValidateLength(t, 3, s.Length())

		s.Pop()
		helpers.ValidateLength(t, 2, s.Length())

		s.Pop()
		helpers.ValidateLength(t, 1, s.Length())
	})

	t.Run("should perform unchained pop", func(t *testing.T) {
		s := NewStack[int]()
		s.Push(1).Push(2).Push(3)
		helpers.ValidateLength(t, 3, s.Length())

		s.Pop().Pop()
		helpers.ValidateLength(t, 1, s.Length())
	})
}

func TestEmpty(t *testing.T) {
	t.Run("should return true for empty stack", func(t *testing.T) {
		if !NewStack[int]().Empty() {
			t.Error("expected stack to be empty")
		}
	})

	t.Run("should return false for empty stack", func(t *testing.T) {
		if NewStack[int]().Push(1).Empty() {
			t.Error("expected stack not to be empty")
		}
	})
}

func TestTop(t *testing.T) {
	t.Run("should return error for empty stack", func(t *testing.T) {
		_, e := NewStack[int]().Top()
		if e == nil {
			t.Errorf("expected \"%s\", but got nil", "stack is empty")
		}
	})

	t.Run("should return top value for non-empty stack", func(t *testing.T) {
		s := NewStack[int]().Push(1).Push(2).Push(3)

		v, e := s.Top()
		if v != 3 {
			t.Errorf("wanted %d as value, got %d", 3, v)
		}
		if e != nil {
			t.Errorf("got %s as error, but wanted nil", e)
		}
	})
}
