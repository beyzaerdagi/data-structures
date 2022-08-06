package binarysearchtree

func (tree *BinarySearchTree) DFSPreOrder() []*Node {

	visited := []*Node{}
	curr := tree.Root
	if curr == nil {
		return visited
	}
	return helperPre(curr, visited)
}

func helperPre(root *Node, visited []*Node) []*Node {
	visited = append(visited, root)

	if root.Left != nil {
		visited = helperPre(root.Left, visited)
	}

	if root.Right != nil {
		visited = helperPre(root.Right, visited)
	}

	return visited
}
