package singlylinkedlist

func (list *SinglyLinkedList) Unshift(node *Node) bool {

	if list.Head == nil {
		list.Head = node
		list.Tail = node
	} else {
		node.Next = list.Head
		list.Head = node
	}
	list.Length++
	return true
}
