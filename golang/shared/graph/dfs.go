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

func (f *Graph[T]) findRouteDFS(node, target *Node[T], path []*Node[T]) []*Node[T] {
	if node == nil || node.Visited() {
		return path
	}

	if node == target {
		return append(path, node)
	}

	node.SetVisited(true)
	for _, adj := range node.GetAdjacents() {
		previousPath := len(path)
		path = f.findRouteDFS(adj, target, path)

		if len(path) > previousPath {
			return append(path, node)
		}
	}

	return path
}
