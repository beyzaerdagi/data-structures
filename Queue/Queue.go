package queue

type Node struct {
	Data int32
	Next *Node
}

type Queue struct {
	First *Node
	Last  *Node
	Size  int32
}
