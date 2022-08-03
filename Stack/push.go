package stack

func (s *Stack) Push(node *Node) {

	if s.First == nil {
		s.First = node
		s.Last = node
	} else {
		curr := s.First
		s.First = node
		s.First.Next = curr
	}
	s.Size++
}
