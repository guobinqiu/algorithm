package algorithm

type Stack struct {
	l *List
}

func NewStack() *Stack {
	return &Stack{
		l: NewList(),
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
