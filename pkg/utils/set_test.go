package utils

import (
	"testing"
)

func TestNewSet(t *testing.T) {
	s := NewSet[int]()
	if s.Get(1) {
		t.Errorf("Expected empty set, but 1 is present")
	}
}

func TestAdd(t *testing.T) {
	s := NewSet[int]()
	s.Add(1)
	if !s.Get(1) {
		t.Errorf("Expected 1 to be in set after Add")
	}
	if s.Get(2) {
		t.Errorf("Expected 2 not to be in set")
	}
}

func TestAddMultiple(t *testing.T) {
	s := NewSet[int]()
	s.AddMultiple([]int{1, 2, 3})
	if !s.Get(1) || !s.Get(2) || !s.Get(3) {
		t.Errorf("Expected 1,2,3 to be in set after AddMultiple")
	}
	if s.Get(4) {
		t.Errorf("Expected 4 not to be in set")
	}
}

func TestGet(t *testing.T) {
	s := NewSet[string]()
	s.Add("hello")
	if !s.Get("hello") {
		t.Errorf("Expected 'hello' to be in set")
	}
	if s.Get("world") {
		t.Errorf("Expected 'world' not to be in set")
	}
}

func TestDelete(t *testing.T) {
	s := NewSet[int]()
	s.Add(42)
	if !s.Get(42) {
		t.Errorf("Expected 42 to be in set")
	}
	s.Delete(42)
	if s.Get(42) {
		t.Errorf("Expected 42 not to be in set after Delete")
	}
}

func TestUnion(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)

	s2 := NewSet[int]()
	s2.Add(2)
	s2.Add(3)

	s1.Union(s2)

	if !s1.Get(1) || !s1.Get(2) || !s1.Get(3) {
		t.Errorf("Expected s1 to contain 1,2,3 after Union")
	}
}

func TestSub(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)
	s1.Add(3)

	s2 := NewSet[int]()
	s2.Add(2)

	s1.Sub(s2)

	if !s1.Get(1) || s1.Get(2) || !s1.Get(3) {
		t.Errorf("Expected s1 to contain 1 and 3, but not 2 after Sub")
	}
}

func TestIntersect(t *testing.T) {
	s1 := NewSet[int]()
	s1.Add(1)
	s1.Add(2)

	s2 := NewSet[int]()
	s2.Add(2)
	s2.Add(3)

	result := Intersect(s1, s2)

	// Intersection should be {2}
	if !result.Get(2) {
		t.Errorf("Expected 2 to be in intersection")
	}
	if result.Get(1) || result.Get(3) {
		t.Errorf("Expected 1 and 3 not to be in intersection")
	}
}