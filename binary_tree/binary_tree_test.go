package binary_tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBinaryTree(t *testing.T) {
	N := 9
	tree := make([]Node, N)

	for i := 0; i < N; i++ {
		tree[i] = Node{parent: -1, left: -1, right: -1, depth: -1, height: -1}
	}

	nodeValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8}
	childrens := [][]int{{1, 4}, {2, 3}, {-1, -1}, {-1, -1}, {5, 8}, {6, 7}, {-1, -1}, {-1, -1}, {-1, -1}}

	for i := 0; i < N; i++ {
		nodeValue := nodeValues[i]
		left := childrens[i][0]
		right := childrens[i][1]

		tree = AppendNode(tree, nodeValue, left, right)
	}

	rootIndex := FindRootNodeIndex(tree)
	assert.Equal(t, 0, rootIndex)

	CalcDepth(tree, 0, 0)
	CalcHeight(tree, 0)

	assert.Equal(t, -1, tree[0].parent)
	assert.Equal(t, 1, tree[0].left)
	assert.Equal(t, 4, tree[0].right)
	assert.Equal(t, 0, tree[0].depth)
	assert.Equal(t, 3, tree[0].height)

	assert.Equal(t, 0, tree[1].parent)
	assert.Equal(t, 2, tree[1].left)
	assert.Equal(t, 3, tree[1].right)
	assert.Equal(t, 1, tree[1].depth)
	assert.Equal(t, 1, tree[1].height)

	assert.Equal(t, 1, tree[2].parent)
	assert.Equal(t, -1, tree[2].left)
	assert.Equal(t, -1, tree[2].right)
	assert.Equal(t, 2, tree[2].depth)
	assert.Equal(t, 0, tree[2].height)

	assert.Equal(t, 1, tree[3].parent)
	assert.Equal(t, -1, tree[3].left)
	assert.Equal(t, -1, tree[3].right)
	assert.Equal(t, 2, tree[3].depth)
	assert.Equal(t, 0, tree[3].height)

	assert.Equal(t, 0, tree[4].parent)
	assert.Equal(t, 5, tree[4].left)
	assert.Equal(t, 8, tree[4].right)
	assert.Equal(t, 1, tree[4].depth)
	assert.Equal(t, 2, tree[4].height)

	assert.Equal(t, 4, tree[5].parent)
	assert.Equal(t, 6, tree[5].left)
	assert.Equal(t, 7, tree[5].right)
	assert.Equal(t, 2, tree[5].depth)
	assert.Equal(t, 1, tree[5].height)

	assert.Equal(t, 5, tree[6].parent)
	assert.Equal(t, -1, tree[6].left)
	assert.Equal(t, -1, tree[6].right)
	assert.Equal(t, 3, tree[6].depth)
	assert.Equal(t, 0, tree[6].height)

	assert.Equal(t, 5, tree[7].parent)
	assert.Equal(t, -1, tree[7].left)
	assert.Equal(t, -1, tree[7].right)
	assert.Equal(t, 3, tree[7].depth)
	assert.Equal(t, 0, tree[7].height)

	assert.Equal(t, 4, tree[8].parent)
	assert.Equal(t, -1, tree[8].left)
	assert.Equal(t, -1, tree[8].right)
	assert.Equal(t, 2, tree[8].depth)
	assert.Equal(t, 0, tree[8].height)
}

func TestDistortedBinaryTree(t *testing.T) {
	N := 7
	tree := make([]Node, N)

	for i := 0; i < N; i++ {
		tree[i] = Node{parent: -1, left: -1, right: -1, depth: -1, height: -1}
	}

	nodeValues := []int{0, 1, 4, 2, 3, 6, 5}
	childrens := [][]int{{-1, -1}, {-1, 4}, {5, -1}, {3, 1}, {-1, 0}, {-1, -1}, {6, -1}}

	for i := 0; i < N; i++ {
		nodeValue := nodeValues[i]
		left := childrens[i][0]
		right := childrens[i][1]

		tree = AppendNode(tree, nodeValue, left, right)
	}

	rootIndex := FindRootNodeIndex(tree)
	assert.Equal(t, 2, rootIndex)

	CalcDepth(tree, rootIndex, 0)
	CalcHeight(tree, rootIndex)

	assert.Equal(t, 3, tree[0].parent)
	assert.Equal(t, -1, tree[0].left)
	assert.Equal(t, -1, tree[0].right)
	assert.Equal(t, 2, tree[0].depth)
	assert.Equal(t, 0, tree[0].height)

	assert.Equal(t, 2, tree[1].parent)
	assert.Equal(t, -1, tree[1].left)
	assert.Equal(t, 4, tree[1].right)
	assert.Equal(t, 1, tree[1].depth)
	assert.Equal(t, 3, tree[1].height)

	assert.Equal(t, -1, tree[2].parent)
	assert.Equal(t, 3, tree[2].left)
	assert.Equal(t, 1, tree[2].right)
	assert.Equal(t, 0, tree[2].depth)
	assert.Equal(t, 4, tree[2].height)

	assert.Equal(t, 2, tree[3].parent)
	assert.Equal(t, -1, tree[3].left)
	assert.Equal(t, 0, tree[3].right)
	assert.Equal(t, 1, tree[3].depth)
	assert.Equal(t, 1, tree[3].height)

	assert.Equal(t, 1, tree[4].parent)
	assert.Equal(t, 5, tree[4].left)
	assert.Equal(t, -1, tree[4].right)
	assert.Equal(t, 2, tree[4].depth)
	assert.Equal(t, 2, tree[4].height)

	assert.Equal(t, 4, tree[5].parent)
	assert.Equal(t, 6, tree[5].left)
	assert.Equal(t, -1, tree[5].right)
	assert.Equal(t, 3, tree[5].depth)
	assert.Equal(t, 1, tree[5].height)

	assert.Equal(t, 5, tree[6].parent)
	assert.Equal(t, -1, tree[6].left)
	assert.Equal(t, -1, tree[6].right)
	assert.Equal(t, 4, tree[6].depth)
	assert.Equal(t, 0, tree[6].height)
}
