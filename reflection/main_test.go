package main

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Name string
	}{expected}

	Walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("wrong number of function calls, got %d want %d", len(got), 1)
	}
}
