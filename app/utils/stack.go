package utils

type Stack struct {
	arr []*int
	currPtr int
	size int
}


func NewStack(size int) *Stack {
	a := make([]*int, size)

	return &Stack{
		arr: a,
		currPtr: -1,
		size: size,
	}
}

func (s *Stack) Pop() *int {
	if s.currPtr < 0 {
		return nil
	}

	res := s.arr[s.currPtr]
	s.arr[s.currPtr] = nil
	s.currPtr = ((s.currPtr - 1) + s.size) % s.size

	return res
}

func (s *Stack) Push(n int) {
	s.currPtr = (s.currPtr + 1) % s.size
	refObj := n
	s.arr[s.currPtr] = &refObj
}




