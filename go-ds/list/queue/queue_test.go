package queue

import (
	"github.com/mddg/go-pro/go-ds/helpers"
	"testing"
)

func ValidateFirst(t *testing.T, q *Queue[int], want int) {
	if q.start.Value != want {
		t.Errorf("wanted %d as first value, got %d", q.start.Value, want)
	}
}

func ValidateLast(t *testing.T, q *Queue[int], want int) {
	if q.end.Value != want {
		t.Errorf("wanted %d as last value, got %d", q.end.Value, want)
	}
}

func TestEnqueue(t *testing.T) {
	t.Run("should perform unchained enqueue", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1)
		q.Enqueue(2)
		q.Enqueue(3)

		ValidateFirst(t, q, 1)
		ValidateLast(t, q, 3)
		helpers.ValidateLength(t, 3, q.length)
	})

	t.Run("should perform chained enqueue", func(t *testing.T) {
		q := NewQueue[int]()

		q.Enqueue(1).Enqueue(2).Enqueue(3)

		ValidateFirst(t, q, 1)
		ValidateLast(t, q, 3)
		helpers.ValidateLength(t, 3, q.length)
	})
}

func TestDequeue(t *testing.T) {
	t.Run("should do nothing if empty", func(t *testing.T) {
		q := NewQueue[int]()

		q.Dequeue()
		helpers.ValidateLength(t, 0, q.length)
	})

	t.Run("should perform unchained dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1).Enqueue(2).Enqueue(3)

		q.Dequeue()
		ValidateFirst(t, q, 2)
		ValidateLast(t, q, 3)
		helpers.ValidateLength(t, 2, q.length)

		q.Dequeue()
		ValidateFirst(t, q, 3)
		ValidateLast(t, q, 3)
		helpers.ValidateLength(t, 1, q.length)
	})

	t.Run("should perform chained dequeue", func(t *testing.T) {
		q := NewQueue[int]()
		q.Enqueue(1).Enqueue(2).Enqueue(3)

		q.Dequeue()
		ValidateFirst(t, q, 2)
		ValidateLast(t, q, 3)
		helpers.ValidateLength(t, 2, q.length)

		q.Dequeue().Dequeue()
		helpers.ValidateLength(t, 0, q.length)
	})
}

func TestEmpty(t *testing.T) {
	t.Run("should return true if queue is empty", func(t *testing.T) {
		if !NewQueue[int]().Empty() {
			t.Error("expected queue to be empty")
		}
	})

	t.Run("should return false if queue is not empty", func(t *testing.T) {
		if NewQueue[int]().Enqueue(1).Empty() {
			t.Error("expected queue not to be empty")
		}
	})
}

func TestEach(t *testing.T) {
	t.Run("should apply the function to each element", func(t *testing.T) {
		res := []int{0, 2, 4}
		q := NewQueue[int]().Enqueue(0).Enqueue(1).Enqueue(2)

		q.Each(func(n *QNode[int], _ int) {
			n.Value = n.Value * 2
		})
		q.Each(func(n *QNode[int], i int) {
			if i*2 != n.Value {
				t.Errorf("wanted %d @ position %d, got %d", res[i], i, n.Value)
			}
		})
	})
}

func TestFirst(t *testing.T) {
	t.Run("should return the first item", func(t *testing.T) {
		q := NewQueue[int]().Enqueue(1)

		v, e := q.First()
		ValidateFirst(t, q, v)
		if e != nil {
			t.Errorf("should have returned error, got %s", e)
		}
	})

	t.Run("should return an error", func(t *testing.T) {
		q := NewQueue[int]()

		_, e := q.First()
		if e == nil {
			t.Errorf("should have not returned an error, got %s", e)
		}
	})
}

func TestLast(t *testing.T) {
	t.Run("should return the last item", func(t *testing.T) {
		q := NewQueue[int]().Enqueue(1).Enqueue(2)

		v, e := q.Last()
		ValidateLast(t, q, v)
		if e != nil {
			t.Errorf("should have returned error, got %s", e)
		}
	})

	t.Run("should return an error", func(t *testing.T) {
		q := NewQueue[int]()

		_, e := q.Last()
		if e == nil {
			t.Errorf("should have not returned an error, got %s", e)
		}
	})
}
