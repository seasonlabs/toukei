package main

import(
	"testing"
)

func TestSanitize(t *testing.T) {
	foo := []byte("   12345   ")

	_, err := sanitize(foo)
	if err != nil {
		t.Error(err)
	}
}

func TestCountCommits(t *testing.T) {
	foo := countCommits()

	if foo == 0 {
		t.Fail()
	} else {
		t.Log(foo)
	}
}