package binarysearchtree

func (tree *BinarySearchTree) DFSInOrder() []*Node {

	visited := []*Node{}
	curr := tree.Root
	if curr == nil {
		return visited
	}
	return helperIn(curr, visited)
}

func helperIn(root *Node, visited []*Node) []*Node {

	if root.Left != nil {
		visited = helperIn(root.Left, visited)
	}
	visited = append(visited, root)
	if root.Right != nil {
		visited = helperIn(root.Right, visited)
	}
	return visited
}
