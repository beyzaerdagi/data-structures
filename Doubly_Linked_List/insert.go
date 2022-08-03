package doublylinkedlist

func (list *DoublyLinkedList) Insert(node *Node, idx int32) bool {

	if idx < 0 || idx > list.Length {
		return false
	}
	if idx == list.Length {
		return list.Push(node)
	}
	if idx == 0 {
		return list.Unshift(node)
	}
	beforeNode := list.Get(idx - 1)
	afterNode := beforeNode.Next
	beforeNode.Next = node
	node.Prev = beforeNode
	node.Next = afterNode
	afterNode.Prev = node
	list.Length++
	return true
}
