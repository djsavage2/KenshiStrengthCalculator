package main

import "testing"

func TestCalculateReqStrength(t *testing.T) {
	// Edge Type 3 Iron Club, expected value retrieved from community wiki
	result := CalculateReqStrength(1.36, 27)
	if result != 55 {
		t.Errorf("Result was incorrect, got: %d, want: %d", result, 55)
	}
}
