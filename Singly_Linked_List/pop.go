package singlylinkedlist

func (list *SinglyLinkedList) Pop() bool {

	if list.Head == nil {
		return false
	}
	curr := list.Head
	newTail := curr
	for curr.Next != nil {
		newTail = curr
		curr = curr.Next
	}
	list.Tail = newTail
	list.Tail.Next = nil
	list.Length--
	if list.Length == 0 {
		list.Head = nil
		list.Tail = nil
	}
	return true
}
