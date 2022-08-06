package binarysearchtree

func (tree *BinarySearchTree) DFSPostOrder() []*Node {

	visited := []*Node{}
	curr := tree.Root
	if curr == nil {
		return visited
	}
	return helperPost(curr, visited)
}

func helperPost(root *Node, visited []*Node) []*Node {

	if root.Left != nil {
		visited = helperPost(root.Left, visited)
	}

	if root.Right != nil {
		visited = helperPost(root.Right, visited)
	}

	visited = append(visited, root)
	return visited
}
