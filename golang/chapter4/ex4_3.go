package chapter4

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/tree"
)

// RunChapter4Exercise3
// List of Depths: Given a binary tree, design an algorithm which creates a
// linked list of all the nodes at each depth (e.g., if you have a tree with
// depth D, you'll have D linked lists).
func RunChapter4Exercise3() {
	fmt.Println("Chapter 4 - Exercise 3: List of Depths")

	bst := tree.BuildSearchTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7})
	if bst == nil {
		fmt.Println("empty tree")
		return
	}

	fmt.Println("Tree:")
	bst.Print()

	fmt.Println("\n--- DFS (pre-order) ---")
	listsDFS := bst.GetNodesByDepthWithDFS()
	for idx, list := range listsDFS {
		fmt.Printf("Depth %d: ", idx)
		list.Print()
	}

	fmt.Println("\n--- BFS (level-order) ---")
	listsBFS := bst.GetNodesByDepthWithBFS()
	for idx, list := range listsBFS {
		fmt.Printf("Depth %d: ", idx)
		list.Print()
	}
}
