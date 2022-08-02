package singlylinkedlist

func (list *SinglyLinkedList) Remove(idx int32) bool {

	if idx < 0 || idx >= list.Length {
		return false
	}
	if idx == list.Length-1 {
		return list.Pop()
	}
	if idx == 0 {
		return list.Shift()
	}
	prevNode := list.Get(idx - 1)
	removeNode := prevNode.Next
	prevNode.Next = removeNode.Next
	list.Length--
	return true
}
