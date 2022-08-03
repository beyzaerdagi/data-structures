package stack

func (s *Stack) Pop() {

	if s.First == nil {
		return
	}
	if s.Size == 1 {
		s.First = nil
		s.Last = nil
	} else {
		s.First = s.First.Next
	}
	s.Size--
}
