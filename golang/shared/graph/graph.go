package graph

import "fmt"

type Graph[T comparable] struct {
	nodes []*Node[T]
}

func NewGraph[T comparable](nodes []*Node[T]) *Graph[T] {
	g := &Graph[T]{
		nodes: make([]*Node[T], 0),
	}

	if nodes != nil {
		g.nodes = nodes
	}

	return g
}

func (g *Graph[T]) IsEmpty() bool {
	return len(g.nodes) == 0
}

func (g *Graph[T]) GetFirst() *Node[T] {
	return g.nodes[0]
}

func (g *Graph[T]) SearchDFS(value T) *Node[T] {
	g.resetNodes()

	first := g.GetFirst()
	return g.dfs(first, value)
}

func (g *Graph[T]) SearchBFS(value T) *Node[T] {
	g.resetNodes()

	first := g.GetFirst()
	return g.bfs(first, value)
}

func (g *Graph[T]) resetNodes() {
	for _, n := range g.nodes {
		n.SetVisited(false)
	}
}

func (g *Graph[T]) FindRouteDFS(a, b *Node[T]) []*Node[T] {
	g.resetNodes()

	if a == nil || b == nil {
		return []*Node[T]{}
	}

	if a == b {
		return []*Node[T]{a}
	}

	path := []*Node[T]{}
	path = g.findRouteDFS(a, b, path)

	// Invert the path
	for i, j := 0, len(path)-1; i < j; i, j = i+1, j-1 {
		path[i], path[j] = path[j], path[i]
	}
	return path
}

func (g *Graph[T]) FindRouteBFS(a, b *Node[T]) []*Node[T] {
	g.resetNodes()

	if a == nil || b == nil {
		return []*Node[T]{}
	}

	if a == b {
		return []*Node[T]{a}
	}

	return g.findRouteBFS(a, b)
}

func (g *Graph[T]) PrintGraph() {
	for _, n := range g.nodes {
		adjs := n.GetAdjacents()
		vals := make([]T, len(adjs))
		for i, a := range adjs {
			vals[i] = a.GetValue()
		}
		fmt.Printf("  %v: %v\n", n.GetValue(), vals)
	}
}
