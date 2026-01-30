package chapter4

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/graph"
)

// RunChapter4Exercise1
// Route Between Nodes: Given a directed graph, design an algorithm to find out
// whether there is a route between two nodes.
func RunChapter4Exercise1() {
	fmt.Println("Chapter 4 - Exercise 1: Route Between Nodes")

	// Graph (adjacency list):
	// 1: [2, 4]
	// 2: [3, 5]
	// 3: []
	// 4: [5]
	// 5: []
	n1 := graph.NewNode(1)
	n2 := graph.NewNode(2)
	n3 := graph.NewNode(3)
	n4 := graph.NewNode(4)
	n5 := graph.NewNode(5)

	n1.AddAdjacent(n2)
	n1.AddAdjacent(n4)
	n2.AddAdjacent(n3)
	n2.AddAdjacent(n5)
	n4.AddAdjacent(n5)

	nodes := []*graph.Node[int]{n1, n2, n3, n4, n5}
	g := graph.NewGraph(nodes)

	fmt.Println("Graph (adjacency list):")
	g.PrintGraph()
	fmt.Println()

	// Find route from 1 to 5
	path := g.FindRouteDFS(n1, n5)
	graph.PrintRoute("1 → 5 (DFS)", path)
	path = g.FindRouteBFS(n1, n5)
	graph.PrintRoute("1 → 5 (BFS)", path)

	// Find route from 1 to 3
	path = g.FindRouteDFS(n1, n3)
	graph.PrintRoute("1 → 3 (DFS)", path)
	path = g.FindRouteBFS(n1, n3)
	graph.PrintRoute("1 → 3 (BFS)", path)

	// No route: 3 to 1
	path = g.FindRouteDFS(n3, n1)
	graph.PrintRoute("3 → 1 (DFS)", path)
	path = g.FindRouteBFS(n3, n1)
	graph.PrintRoute("3 → 1 (BFS)", path)
}
