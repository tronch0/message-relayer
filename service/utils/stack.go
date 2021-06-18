package utils

import (
	model2 "message-relayer/service/model"
)

// Stack implementation with circular array under the hood
type Stack struct {
	arr []*model2.Message
	currPtr int
	size int
}


func NewStack(size int) *Stack {
	a := make([]*model2.Message, size)

	return &Stack{
		arr: a,
		currPtr: -1,
		size: size,
	}
}

func (s *Stack) Pop() *model2.Message {
	if s.currPtr < 0 {
		return nil
	}

	res := s.arr[s.currPtr]
	s.arr[s.currPtr] = nil
	s.currPtr = ((s.currPtr - 1) + s.size) % s.size

	return res
}

func (s *Stack) Push(msg model2.Message) {
	s.currPtr = (s.currPtr + 1) % s.size
	s.arr[s.currPtr] = &msg
}




