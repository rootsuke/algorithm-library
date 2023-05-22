package binary_tree

type Node struct {
	parent int
	left   int
	right  int
	depth  int
	height int
}

func CalcDepth(tree []Node, nodeValue int, depth int) {
	tree[nodeValue].depth = depth

	if tree[nodeValue].left != -1 {
		CalcDepth(tree, tree[nodeValue].left, depth+1)
	}

	if tree[nodeValue].right != -1 {
		CalcDepth(tree, tree[nodeValue].right, depth+1)
	}
}

func CalcHeight(tree []Node, nodeValue int) int {
	currentNode := &tree[nodeValue]
	leftHeight := 0
	rightHeight := 0

	if currentNode.left != -1 {
		leftHeight = CalcHeight(tree, currentNode.left) + 1
	}

	if currentNode.right != -1 {
		rightHeight = CalcHeight(tree, currentNode.right) + 1
	}

	if leftHeight >= rightHeight {
		currentNode.height = leftHeight
		return leftHeight
	} else {
		currentNode.height = rightHeight
		return rightHeight
	}
}

func AppendNode(tree []Node, nodeValue int, left int, right int) []Node {
	if left != -1 {
		tree[nodeValue].left = left
		tree[left].parent = nodeValue
	}

	if right != -1 {
		tree[nodeValue].right = right
		tree[right].parent = nodeValue
	}

	return tree
}

func FindRootNodeIndex(tree []Node) int {
	rootNodeIndex := 0
	for i := 0; i < len(tree); i++ {
		if tree[i].parent == -1 {
			rootNodeIndex = i
			break
		}
	}

	return rootNodeIndex
}
