package main

import (
	"testing"
)

func TestFactorialNumber(t *testing.T) {
	// check for error
	want := 0
	got, err := FactorialNumber(-10)
	if err == nil {
		t.Errorf("FactorialNumber() = %v, want %v", got, err)
	}

	want = 1
	got, err = FactorialNumber(0)
	if err == nil && got != uint64(want) {
		t.Errorf("Expected '%v' but got '%v'", 1, got)
	}

	want = 40320
	got, err = FactorialNumber(8)
	if err == nil && got != uint64(want) {
		t.Errorf("FactorialNumber() = %v, want %v", got, want)
	}
}
