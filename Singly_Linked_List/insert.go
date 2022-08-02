package singlylinkedlist

func (list *SinglyLinkedList) Insert(node *Node, idx int32) bool {

	if idx < 0 || idx > list.Length {
		return false
	}
	if idx == list.Length {
		return list.Push(node)
	}
	if idx == 0 {
		return list.Unshift(node)
	}
	prevNode := list.Get(idx - 1)
	tmp := prevNode.Next
	prevNode.Next = node
	node.Next = tmp
	list.Length++
	return true
}
