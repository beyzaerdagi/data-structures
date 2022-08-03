package doublylinkedlist

func (list *DoublyLinkedList) Push(node *Node) bool {

	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Tail.Next = node
		node.Prev = list.Tail
		list.Tail = node
	}
	list.Length++
	return true
}
