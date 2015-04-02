package main_test

import "testing"

func TestHandlesNonExistingItems(t *testing.T) {
	if x := 1; x == 1 {
		t.Errorf("expected 1, got %d", x)
	}
}
