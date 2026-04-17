package main

import (
	"testing"
)

func TestListNumbers(t *testing.T) {
	want := "112,119,126,133,147,154,161,168,182,189,196"
	got := ListNumbers(100, 200)

	if got != want {
		t.Errorf("ListNumbers() = %v, want %v", got, want)
	}
}
