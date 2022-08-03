package doublylinkedlist

func (list *DoublyLinkedList) Get(idx int32) *Node {

	if idx >= list.Length || idx < 0 {
		return nil
	}
	count := 0
	curr := list.Head
	if idx <= list.Length/2 {
		for count != int(idx) {
			curr = curr.Next
			count++
		}
	} else {
		count = int(list.Length) - 1
		curr = list.Tail
		for count != int(idx) {
			curr = curr.Prev
			count--
		}
	}
	return curr
}
