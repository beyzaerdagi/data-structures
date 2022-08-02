package singlylinkedlist

type Node struct {
	Data int32
	Next *Node
}

type SinglyLinkedList struct {
	Head   *Node
	Tail   *Node
	Length int32
}
