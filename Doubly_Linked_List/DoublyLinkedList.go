package doublylinkedlist

type Node struct {
	Data int32
	Next *Node
	Prev *Node
}

type DoublyLinkedList struct {
	Length int32
	Head   *Node
	Tail   *Node
}
