package singlylinkedlist

func (list *SinglyLinkedList) Set(newNode *Node, idx int32) bool {

	currNode := list.Get(idx)

	if currNode != nil {
		currNode = newNode
		return true
	}
	return false
}
