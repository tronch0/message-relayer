package utils

import (
	"testing"
)

func TestEmptyStack(t *testing.T) {
	s := NewStack(3)

	for i := 0; i < 3; i++ {
		res := s.Pop()
		if res != nil {
			t.Fatalf("res value failed, expected val: nil, actual val: %v", res)
		}
	}
}

func TestHappyFlow(t *testing.T) {
	s := NewStack(3)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	res := s.Pop()
	if *res != 3 {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 3,*res)
	}
	res = s.Pop()
	if *res != 2 {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 2,*res)
	}

	res = s.Pop()
	if *res != 1 {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 1,*res)
	}

	res = s.Pop()
	if res != nil {
		t.Fatalf("res value failed, expected val: nil, actual val: %d",res)
	}
}

func TestOverwriteValues(t *testing.T) {
	s := NewStack(2)
	s.Push(1)
	s.Push(2)
	s.Push(3)

	res := s.Pop()
	if *res != 3 {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 3,*res)
	}

	res = s.Pop()
	if *res != 2 {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 2,*res)
	}

	res = s.Pop()
	if res != nil {
		t.Fatalf("res value failed, expected val: nil, actual val: %d",res)
	}
}