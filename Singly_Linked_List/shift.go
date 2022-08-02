package singlylinkedlist

func (list *SinglyLinkedList) Shift() bool {

	if list.Head == nil {
		return false
	}
	list.Head = list.Head.Next
	list.Length--
	if list.Length == 0 {
		list.Tail = nil
	}
	return true
}
