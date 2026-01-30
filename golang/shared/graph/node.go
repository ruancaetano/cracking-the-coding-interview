package graph

type Node[T comparable] struct {
	value     T
	adjacents []*Node[T]
	visited   bool
}

func NewNode[T comparable](v T) *Node[T] {
	return &Node[T]{
		value:     v,
		adjacents: make([]*Node[T], 0),
		visited:   false,
	}
}

func (n *Node[T]) GetValue() T {
	return n.value
}

func (n *Node[T]) Visited() bool {
	return n.visited
}

func (n *Node[T]) SetVisited(v bool) {
	n.visited = v
}

func (n *Node[T]) GetAdjacents() []*Node[T] {
	return n.adjacents
}

func (n *Node[T]) AddAdjacent(node *Node[T]) {
	n.adjacents = append(n.adjacents, node)
}
