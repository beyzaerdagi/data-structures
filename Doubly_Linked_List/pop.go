package doublylinkedlist

func (list *DoublyLinkedList) Pop() bool {

	if list.Head == nil {
		return false
	}
	if list.Length == 1 {
		list.Head = nil
		list.Tail = nil
	} else {
		poppedNode := list.Tail
		list.Tail = poppedNode.Prev
		list.Tail.Next = nil
		poppedNode.Prev = nil
	}
	list.Length--
	return true
}
