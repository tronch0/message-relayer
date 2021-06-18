package utils

import (
	"message-relayer/model"
	"message-relayer/model/messagetype"
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
	s.Push(model.Message{Type: messagetype.StartNewRound, Data: []byte("1")})
	s.Push(model.Message{Type: messagetype.ReceivedAnswer, Data: []byte("2")})
	s.Push(model.Message{Type: messagetype.Undefined, Data: []byte("3")})

	res := s.Pop()
	if res.Type != messagetype.Undefined {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 3,res.Type)
	}
	res = s.Pop()
	if res.Type != messagetype.ReceivedAnswer {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 2,res.Type)
	}

	res = s.Pop()
	if res.Type != messagetype.StartNewRound {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 1,res.Type)
	}

	res = s.Pop()
	if res != nil {
		t.Fatalf("res value failed, expected val: nil, actual val: %d",res)
	}
}

func TestOverwriteValues(t *testing.T) {
	s := NewStack(2)
	s.Push(model.Message{Type: messagetype.StartNewRound, Data: []byte("1")})
	s.Push(model.Message{Type: messagetype.ReceivedAnswer, Data: []byte("2")})
	s.Push(model.Message{Type: messagetype.Undefined, Data: []byte("3")})

	res := s.Pop()
	if res.Type != messagetype.Undefined {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 3,*res)
	}

	res = s.Pop()
	if res.Type != messagetype.ReceivedAnswer {
		t.Fatalf("res value failed, expected val: %d, actual val: %d", 2,*res)
	}

	res = s.Pop()
	if res != nil {
		t.Fatalf("res value failed, expected val: nil, actual val: %d",res)
	}
}

