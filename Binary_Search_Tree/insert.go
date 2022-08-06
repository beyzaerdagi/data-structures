package binarysearchtree

func (tree *BinarySearchTree) Insert(node *Node) {

	if tree.Root == nil {
		tree.Root = node
		return
	}
	curr := tree.Root
	for {
		if curr.Key == node.Key {
			return
		}
		if curr.Key > node.Key {
			if curr.Left == nil {
				curr.Left = node
				return
			}
			curr = curr.Left
		} else {
			if curr.Right == nil {
				curr.Right = node
				return
			}
			curr = curr.Right
		}
	}
}
