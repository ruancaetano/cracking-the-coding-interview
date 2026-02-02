package tree

import (
	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/queue"
)

func (b *BinarySearchTree[T]) getNodesByDepthWithDFS(node *Node[T], lists []*shared.LinkedList[T]) []*shared.LinkedList[T] {
	if node == nil {
		return lists
	}

	if len(lists) < node.Depth+1 {
		lists = append(lists, &shared.LinkedList[T]{})
	}

	listNode := shared.NewNode(node.Value)
	lists[node.Depth].Add(listNode)

	if node.Left != nil {
		node.Left.Depth = node.Depth + 1
		lists = b.getNodesByDepthWithDFS(node.Left, lists)
	}

	if node.Right != nil {
		node.Right.Depth = node.Depth + 1
		lists = b.getNodesByDepthWithDFS(node.Right, lists)
	}

	return lists
}

func (b *BinarySearchTree[T]) getNodesByDepthWithBFS() []*shared.LinkedList[T] {
	lists := make([]*shared.LinkedList[T], 0)

	// todo: add dynamic queue
	nodeQueue := queue.NewQueue[*Node[T]](100)
	if b.Root != nil {
		b.Root.Depth = 0
		nodeQueue.Enqueue(b.Root)
	}

	for !nodeQueue.IsEmpty() {
		node, _ := nodeQueue.Dequeue()

		if node == nil {
			continue
		}

		depth := node.Depth

		if len(lists) < depth+1 {
			lists = append(lists, &shared.LinkedList[T]{})
		}

		listNode := shared.NewNode(node.Value)
		lists[depth].Add(listNode)

		if node.Left != nil {
			node.Left.Depth = node.Depth + 1
			nodeQueue.Enqueue(node.Left)
		}

		if node.Right != nil {
			node.Right.Depth = node.Depth + 1
			nodeQueue.Enqueue(node.Right)
		}
	}

	return lists
}
