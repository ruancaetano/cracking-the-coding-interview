package tree

import (
	"cmp"
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

type BinarySearchTree[T cmp.Ordered] struct {
	Root *Node[T]
}

func (b *BinarySearchTree[T]) GetNodesByDepthWithDFS() []*shared.LinkedList[T] {
	lists := make([]*shared.LinkedList[T], 0)
	if b.Root != nil {
		b.Root.Depth = 0
	}
	return b.getNodesByDepthWithDFS(b.Root, lists)
}

func (b *BinarySearchTree[T]) GetNodesByDepthWithBFS() []*shared.LinkedList[T] {
	return b.getNodesByDepthWithBFS()
}

func (b *BinarySearchTree[T]) Print() {
	b.printStruct(b.Root, "", true, true)
}

func (b *BinarySearchTree[T]) printStruct(n *Node[T], prefix string, isTail bool, isRoot bool) {
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

	// Left and right as "children" (left first, then right)
	hasLeft, hasRight := n.Left != nil, n.Right != nil
	if hasLeft && hasRight {
		b.printStruct(n.Left, childPrefix, false, false)
		b.printStruct(n.Right, childPrefix, true, false)
	} else if hasLeft {
		b.printStruct(n.Left, childPrefix, true, false)
	} else if hasRight {
		b.printStruct(n.Right, childPrefix, true, false)
	}
}
