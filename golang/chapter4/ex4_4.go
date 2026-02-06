package chapter4

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/tree"
)

// RunChapter4Exercise4
// Check Balanced: Implement a function to check if a binary tree is balanced.
// For the purposes of this question, a balanced tree is defined to be a tree
// such that the heights of the two subtrees of any node never differ by more
// than one.
func RunChapter4Exercise4() {
	fmt.Println("Chapter 4 - Exercise 4: Check Balanced")

	balanced := tree.FromAdjacentList(1, map[int][]*int{
		1: {tree.Ptr(2), tree.Ptr(3)},
		2: {tree.Ptr(4), tree.Ptr(5)},
		3: {tree.Ptr(6), tree.Ptr(7)},
	})
	fmt.Println("Balanced tree:")
	balanced.Print()
	fmt.Printf("IsBalanced: %v\n\n", balanced.IsBalanced())

	unbalanced := tree.FromAdjacentList(1, map[int][]*int{
		1: {tree.Ptr(2), tree.Ptr(3)},
		2: {tree.Ptr(4), tree.Ptr(5)},
		4: {tree.Ptr(6), nil},
		6: {tree.Ptr(7), nil},
	})
	fmt.Println("Unbalanced tree (chain on left):")
	unbalanced.Print()
	fmt.Printf("IsBalanced: %v\n\n", unbalanced.IsBalanced())

	// Tricky: balanced at root (both sides height 3) but unbalanced at nodes 2 and 3
	tricky := tree.FromAdjacentList(1, map[int][]*int{
		1: {tree.Ptr(2), tree.Ptr(3)},
		2: {tree.Ptr(4), nil},
		3: {nil, tree.Ptr(5)},
		4: {tree.Ptr(6), nil},
		5: {nil, tree.Ptr(7)},
	})
	fmt.Println("Tricky tree (balanced at root, unbalanced at nodes 2 and 3):")
	tricky.Print()
	fmt.Printf("IsBalanced: %v\n", tricky.IsBalanced())
}
