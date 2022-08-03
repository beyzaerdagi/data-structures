package doublylinkedlist

func (list *DoublyLinkedList) Shift() bool {

	if list.Head == nil {
		return false
	}
	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil
	} else {
		shiftedNode := list.Head
		list.Head = shiftedNode.Next
		list.Head.Prev = nil
		shiftedNode.Next = nil
	}
	list.Length--
	return true
}
