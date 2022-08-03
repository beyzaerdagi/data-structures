package stack

type Node struct {
	Data int32
	Next *Node
}

type Stack struct {
	First *Node
	Last  *Node
	Size  int32
}
