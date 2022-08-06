package binarysearchtree

type Node struct {
	Key   int32
	Left  *Node
	Right *Node
}

type BinarySearchTree struct {
	Root *Node
}
