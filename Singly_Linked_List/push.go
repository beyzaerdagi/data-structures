package singlylinkedlist

func (list *SinglyLinkedList) Push(node *Node) bool {

	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		list.Tail = node
	}
	list.Length++
	return true
}
