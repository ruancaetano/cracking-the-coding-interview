package tree

import (
	"fmt"
)

type BinaryTree[T any] struct {
	Root *Node[T]
}

// FromAdjacentList builds a tree from an adjacent list. adj[node] = [left, right]; nil means no child.
func FromAdjacentList[T comparable](rootVal T, adj map[T][]*T) *BinaryTree[T] {
	return &BinaryTree[T]{Root: buildFromAdjacentList(rootVal, adj)}
}

// Ptr returns a pointer to v. Use with FromAdjacentList for child values.
func Ptr[T any](v T) *T { return &v }

func (b *BinaryTree[T]) CalcTreeHeight() int {
	return b.calcBranchHeight(b.Root)
}

func (b *BinaryTree[T]) CalcBranchHeight(node *Node[T]) int {
	return b.calcBranchHeight(node)
}

func (b *BinaryTree[T]) IsBalanced() bool {
	return b.isBalanced(b.Root)
}

func (b *BinaryTree[T]) isBalanced(n *Node[T]) bool {
	if n == nil {
		return true
	}
	leftH := b.calcBranchHeight(n.Left)
	rightH := b.calcBranchHeight(n.Right)
	if leftH-rightH > 1 || rightH-leftH > 1 {
		return false
	}
	return b.isBalanced(n.Left) && b.isBalanced(n.Right)
}

func (b *BinaryTree[T]) calcBranchHeight(node *Node[T]) int {
	if node == nil {
		return 0
	}

	leftHeight := 1 + b.calcBranchHeight(node.Left)
	rightHeight := 1 + b.calcBranchHeight(node.Right)

	if leftHeight > rightHeight {
		return leftHeight
	}
	return rightHeight
}

func (b *BinaryTree[T]) Print() {
	b.printStruct(b.Root, "", true, true)
}

func buildFromAdjacentList[T comparable](val T, adj map[T][]*T) *Node[T] {
	n := &Node[T]{Value: val}
	children := adj[val]
	if len(children) >= 1 && children[0] != nil {
		n.Left = buildFromAdjacentList(*children[0], adj)
	}
	if len(children) >= 2 && children[1] != nil {
		n.Right = buildFromAdjacentList(*children[1], adj)
	}
	return n
}

func (b *BinaryTree[T]) printStruct(n *Node[T], prefix string, isTail bool, isRoot bool) {
	if n == nil {
		return
	}

	connector := "├── "
	if isTail {
		connector = "└── "
	}
	if isRoot {
		connector = "    "
	}

	fmt.Println(prefix + connector + fmt.Sprint(n.Value))

	childPrefix := prefix
	if isTail {
		childPrefix += "    "
	} else {
		childPrefix += "│   "
	}

	hasLeft, hasRight := n.Left != nil, n.Right != nil
	if hasLeft {
		b.printStruct(n.Left, childPrefix, !hasRight, false)
	} else if hasRight {
		fmt.Println(childPrefix + "├── nil")
	}

	if hasRight {
		b.printStruct(n.Right, childPrefix, true, false)
	} else if hasLeft {
		fmt.Println(childPrefix + "└── nil")
	}
}
