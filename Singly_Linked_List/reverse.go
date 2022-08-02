package singlylinkedlist

func (list *SinglyLinkedList) Reverse() {

	if list.Head == nil {
		return
	}
	if list.Length == 1 {
		return
	}
	node := list.Head
	list.Head, list.Tail = list.Tail, list.Head
	var next *Node
	var prev *Node
	for i := 0; i < int(list.Length); i++ {
		next = node.Next
		node.Next = prev
		prev = node
		node = next
	}
}
