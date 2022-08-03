package doublylinkedlist

func (list *DoublyLinkedList) Remove(idx int32) bool {

	if idx < 0 || idx >= list.Length {
		return false
	}
	if idx == list.Length-1 {
		return list.Pop()
	}
	if idx == 0 {
		return list.Shift()
	}
	removedNode := list.Get(idx)
	removedNode.Prev.Next = removedNode.Next
	removedNode.Next.Prev = removedNode.Prev
	removedNode.Next = nil
	removedNode.Prev = nil
	list.Length--
	return true
}
