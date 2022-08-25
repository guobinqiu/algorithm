package stack

import "algorithm/list"

type Stack struct {
	l *list.List
}

func New() *Stack {
	return &Stack{
		l: list.New(),
	}
}

func (s *Stack) Pop() interface{} {
	if s.l.Len() > 0 {
		e := s.l.Last()
		s.l.Remove(e)
		return e.Value
	}
	return nil
}

func (s *Stack) Push(value interface{}) {
	s.l.Append(value)
}

func (s *Stack) IsEmpty() bool {
	return s.l.Len() == 0
}

func (s *Stack) Clear() {
	s.l = nil
}

func (s *Stack) Size() int {
	return s.l.Len()
}
