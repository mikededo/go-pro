package helpers

import "testing"

func ValidateLength(t *testing.T, want, got int) {
	if got != want {
		t.Errorf("got %d for length, but wanted %d", got, want)
	}
}
