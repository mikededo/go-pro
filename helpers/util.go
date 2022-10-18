package helpers

import "testing"

func ValidateLength(t *testing.T, wanted, got int) {
	if got != wanted {
		t.Errorf("got %d for length, but wanted %d", got, wanted)
	}
}
