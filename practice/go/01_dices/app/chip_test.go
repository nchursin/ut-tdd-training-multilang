package app

import "testing"

func TestChipEquality(t *testing.T) {
	if !((Chip{Amount: 10}).Equal(Chip{Amount: 10})) {
		t.Fatal("chips with equal amounts must be equal")
	}
}
