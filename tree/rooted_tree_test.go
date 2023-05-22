package tree

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRootedTree(t *testing.T) {
	N := 13
	tree := make([]Node, N)

	for i := 0; i < N; i++ {
		tree[i] = Node{parent: -1, leftChild: -1, rightSibling: -1}
	}

	nodeValues := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12}
	childrens := [][]int{{1, 4, 10}, {2, 3}, {}, {}, {5, 6, 7}, {}, {}, {8, 9}, {}, {}, {11, 12}, {}, {}}
	for i := 0; i < N; i++ {
		nodeValue := nodeValues[i]
		children := childrens[i]

		tree = AppendNode(tree, nodeValue, children)
	}

	rootIndex := FindRootNodeIndex(tree)
	assert.Equal(t, 0, rootIndex)

	CalcDepth(tree, 0, 0)

	assert.Equal(t, -1, tree[0].parent)
	assert.Equal(t, 1, tree[0].leftChild)
	assert.Equal(t, -1, tree[0].rightSibling)
	assert.Equal(t, 0, tree[0].depth)

	assert.Equal(t, 0, tree[1].parent)
	assert.Equal(t, 2, tree[1].leftChild)
	assert.Equal(t, 4, tree[1].rightSibling)
	assert.Equal(t, 1, tree[1].depth)

	assert.Equal(t, 1, tree[2].parent)
	assert.Equal(t, -1, tree[2].leftChild)
	assert.Equal(t, 3, tree[2].rightSibling)
	assert.Equal(t, 2, tree[2].depth)

	assert.Equal(t, 1, tree[3].parent)
	assert.Equal(t, -1, tree[3].leftChild)
	assert.Equal(t, -1, tree[3].rightSibling)
	assert.Equal(t, 2, tree[3].depth)

	assert.Equal(t, 0, tree[4].parent)
	assert.Equal(t, 5, tree[4].leftChild)
	assert.Equal(t, 10, tree[4].rightSibling)
	assert.Equal(t, 1, tree[4].depth)

	assert.Equal(t, 4, tree[5].parent)
	assert.Equal(t, -1, tree[5].leftChild)
	assert.Equal(t, 6, tree[5].rightSibling)
	assert.Equal(t, 2, tree[5].depth)

	assert.Equal(t, 4, tree[6].parent)
	assert.Equal(t, -1, tree[6].leftChild)
	assert.Equal(t, 7, tree[6].rightSibling)
	assert.Equal(t, 2, tree[6].depth)

	assert.Equal(t, 4, tree[7].parent)
	assert.Equal(t, 8, tree[7].leftChild)
	assert.Equal(t, -1, tree[7].rightSibling)
	assert.Equal(t, 2, tree[7].depth)

	assert.Equal(t, 7, tree[8].parent)
	assert.Equal(t, -1, tree[8].leftChild)
	assert.Equal(t, 9, tree[8].rightSibling)
	assert.Equal(t, 3, tree[8].depth)

	assert.Equal(t, 7, tree[9].parent)
	assert.Equal(t, -1, tree[9].leftChild)
	assert.Equal(t, -1, tree[9].rightSibling)
	assert.Equal(t, 3, tree[9].depth)

	assert.Equal(t, 0, tree[10].parent)
	assert.Equal(t, 11, tree[10].leftChild)
	assert.Equal(t, -1, tree[10].rightSibling)
	assert.Equal(t, 1, tree[10].depth)

	assert.Equal(t, 10, tree[11].parent)
	assert.Equal(t, -1, tree[11].leftChild)
	assert.Equal(t, 12, tree[11].rightSibling)
	assert.Equal(t, 2, tree[11].depth)

	assert.Equal(t, 10, tree[12].parent)
	assert.Equal(t, -1, tree[12].leftChild)
	assert.Equal(t, -1, tree[12].rightSibling)
	assert.Equal(t, 2, tree[12].depth)
}
