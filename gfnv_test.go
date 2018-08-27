package gfnv

import "testing"

func TestFnv32(t *testing.T) {
	got1 := Fnv32("hello world")
	got2 := Fnv32("hello world")

	if got1 != got2 {
		t.Errorf("Fnv32 is wrong. expected: %v, got: %v", got1, got2)
	}
}

func TestFnv32a(t *testing.T) {
	got1 := Fnv32a("hello world")
	got2 := Fnv32a("hello world")

	if got1 != got2 {
		t.Errorf("Fnv32a is wrong. expected: %v, got: %v", got1, got2)
	}
}

func TestFnv64(t *testing.T) {
	got1 := Fnv64("hello world")
	got2 := Fnv64("hello world")

	if got1 != got2 {
		t.Errorf("Fnv32 is wrong. expected: %v, got: %v", got1, got2)
	}
}

func TestFnv64a(t *testing.T) {
	got1 := Fnv64a("hello world")
	got2 := Fnv64a("hello world")

	if got1 != got2 {
		t.Errorf("Fnv32a is wrong. expected: %v, got: %v", got1, got2)
	}
}
