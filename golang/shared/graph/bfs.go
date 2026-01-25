package graph

import (
	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/queue"
)

func (g *Graph[T]) bfs(root *Node[T], value T) *Node[T] {
	if root == nil {
		return root
	}

	// TODO: a queue with dynamic allocation is better
	queue := queue.NewQueue[(*Node[T])](1000)

	queue.Enqueue(root)

	for !queue.IsEmpty() {
		node, _ := queue.Dequeue()
		node.SetVisited(true)

		if node.GetValue() == value {
			return node
		}

		for _, n := range node.GetAdjacents() {
			if n.Visited() {
				continue
			}
			queue.Enqueue(n)
		}
	}

	return nil
}
