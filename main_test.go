package main

import (
	"testing"
)

func TestPow2(t *testing.T) {
	want := 4
	get := pow(2, 2)

	if want != get {
		t.Errorf("Error in TestPow. Want %d, Get %d", want, get)
	}
}

func TestPow3(t *testing.T) {
	want := 27
	get := pow(3, 3)

	if want != get {
		t.Errorf("Error in TestPow. Want %d, Get %d", want, get)
	}
}

func TestPow5(t *testing.T) {
	want := 125
	get := pow(5, 3)

	if want != get {
		t.Errorf("Error in TestPow. Want %d, Get %d", want, get)
	}
}

func BenchmarkPow(b *testing.B) {
	for i := 0; i < b.N; i++ {
		pow(5, 5)
	}
}
