package graph

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
