package doublylinkedlist

func (list *DoublyLinkedList) Set(newNode *Node, idx int32) bool {

	currNode := list.Get(idx)

	if currNode != nil {
		currNode = newNode
		return true
	}
	return false
}
