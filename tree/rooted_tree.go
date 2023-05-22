package tree

type Node struct {
	parent       int
	leftChild    int
	rightSibling int
	depth        int
}

func CalcDepth(tree []Node, nodeValue int, depth int) {
	tree[nodeValue].depth = depth

	if tree[nodeValue].leftChild != -1 {
		CalcDepth(tree, tree[nodeValue].leftChild, depth+1)
	}

	if tree[nodeValue].rightSibling != -1 {
		CalcDepth(tree, tree[nodeValue].rightSibling, depth)
	}
}

func AppendNode(tree []Node, nodeValue int, children []int) []Node {
	for j := 0; j < len(children); j++ {
		if j == 0 {
			tree[nodeValue].leftChild = children[0]
		} else {
			tree[children[j-1]].rightSibling = children[j]
		}
		tree[children[j]].parent = nodeValue
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
