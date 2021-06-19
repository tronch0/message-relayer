package utils

import (
	"message-relayer/service/model"
)

// Stack implementation with circular array under the hood
type Stack struct {
	arr []*model.Message
	currPtr int
	size int
}


func NewStack(size int) *Stack {
	a := make([]*model.Message, size)

	return &Stack{
		arr: a,
		currPtr: -1,
		size: size,
	}
}

func (s *Stack) Pop() *model.Message {
	if s.currPtr < 0 {
		return nil
	}

	res := s.arr[s.currPtr]
	s.arr[s.currPtr] = nil
	s.currPtr = ((s.currPtr - 1) + s.size) % s.size

	return res
}

func (s *Stack) Push(msg model.Message) {
	s.currPtr = (s.currPtr + 1) % s.size
	s.arr[s.currPtr] = &msg
}




