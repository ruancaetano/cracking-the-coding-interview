package graph

func (g *Graph[T]) dfs(root *Node[T], value T) *Node[T] {
	if root == nil {
		return nil
	}

	root.SetVisited(true)

	if root.GetValue() == value {
		return root
	}

	for _, adj := range root.GetAdjacents() {
		if adj.Visited() {
			continue
		}

		if result := g.dfs(adj, value); result != nil {
			return result
		}
	}

	return nil
}
