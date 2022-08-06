package binarysearchtree

func (tree *BinarySearchTree) Find(node *Node) bool {

	if tree.Root == nil {
		return false
	}

	curr := tree.Root
	for curr != nil {
		if curr.Key > node.Key {
			curr = curr.Left
		} else if curr.Key < node.Key {
			curr = curr.Right
		} else {
			return true
		}
	}
	return false
}
