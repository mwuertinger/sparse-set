package set

import (
	"testing"
	"math/rand"
)

func TestSet(t *testing.T) {
	max := uint(1024)

	s := New(max)
	randomize(s)
	for i := uint(0); i < max; i++ {
		if s.Contains(i) {
			t.Fatalf("set should not contain %v", i)
		}
	}

	data := []uint{17, 73, 23, 42}
	for _, x := range data {
		s.Add(x)
	}
	for _, x := range data {
		if !s.Contains(x) {
			t.Fatalf("set should contain %v", x)
		}
	}

	s.Add(17)
	s.Add(42)
	s.Clear()
	s.Add(23)

	if s.Contains(17) || s.Contains(42) {
		t.Fatal("set should not contain 17 or 42")
	}
	if !s.Contains(23) {
		t.Fatal("set should contain 23")
	}
}

func randomize(s *Set) {
	for i := range s.dense {
		s.dense[i] = uint(rand.Int31())
		s.sparse[i] = uint(rand.Int31())
	}
}
