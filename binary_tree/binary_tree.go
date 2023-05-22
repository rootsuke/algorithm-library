package binary_tree

import "fmt"

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

func Preorder(tree []Node, id int) {
	fmt.Printf(" %v", id)

	if tree[id].left != -1 {
		Preorder(tree, tree[id].left)
	}

	if tree[id].right != -1 {
		Preorder(tree, tree[id].right)
	}
}

func Inorder(tree []Node, id int) {
	if tree[id].left != -1 {
		Inorder(tree, tree[id].left)
	}

	fmt.Printf(" %v", id)

	if tree[id].right != -1 {
		Inorder(tree, tree[id].right)
	}
}

func Postorder(tree []Node, id int) {
	if tree[id].left != -1 {
		Postorder(tree, tree[id].left)
	}

	if tree[id].right != -1 {
		Postorder(tree, tree[id].right)
	}

	fmt.Printf(" %v", id)
}
