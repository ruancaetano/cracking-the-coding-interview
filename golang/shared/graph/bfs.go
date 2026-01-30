package graph

import (
	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/queue"
)

type bfsNodePath[T comparable] struct {
	node *Node[T]
	path []*Node[T]
}

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

func (g *Graph[T]) findRouteBFS(from, to *Node[T]) []*Node[T] {

	// TODO: a queue with dynamic allocation is better
	queue := queue.NewQueue[*bfsNodePath[T]](1000)

	queue.Enqueue(&bfsNodePath[T]{
		node: from,
		path: []*Node[T]{from},
	})

	for !queue.IsEmpty() {
		nodePath, _ := queue.Dequeue()
		nodePath.node.SetVisited(true)

		if nodePath.node == to {
			return nodePath.path
		}

		for _, adj := range nodePath.node.GetAdjacents() {
			if adj.Visited() {
				continue
			}

			adjPath := &bfsNodePath[T]{
				node: adj,
				path: append(nodePath.path, adj),
			}
			queue.Enqueue(adjPath)
		}
	}

	return []*Node[T]{}
}
