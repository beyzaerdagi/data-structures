package singlylinkedlist

func (list *SinglyLinkedList) Get(idx int32) *Node {

	if idx >= list.Length || idx < 0 {
		return nil
	}
	curr := list.Head
	count := 0
	for count != int(idx) {
		curr = curr.Next
		count++
	}
	return curr
}
