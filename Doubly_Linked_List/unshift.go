package doublylinkedlist

func (list *DoublyLinkedList) Unshift(node *Node) bool {

	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		list.Head.Prev = node
		node.Next = list.Head
		list.Head = node
	}
	list.Length++
	return true
}
