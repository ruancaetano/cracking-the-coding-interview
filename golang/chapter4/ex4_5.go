package chapter4

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/tree"
)

// RunChapter4Exercise5
// Validate BST: Implement a function to check if a binary tree is a binary
// search tree.
func RunChapter4Exercise5() {
	fmt.Println("Chapter 4 - Exercise 5: Validate BST")

	valid := tree.BuildSearchTreeFromSortedArray([]int{1, 2, 3, 4, 5, 6, 7})
	fmt.Println("Valid BST (from sorted array):")
	valid.Print()
	fmt.Printf("IsBinarySearchTree: %v\n\n", valid.IsBinarySearchTree())

	invalid := &tree.BinarySearchTree[int]{
		Root: &tree.Node[int]{
			Value: 10,
			Left:  &tree.Node[int]{Value: 15},
			Right: &tree.Node[int]{Value: 20},
		},
	}
	fmt.Println("Invalid BST (left child 15 > root 10):")
	invalid.Print()
	fmt.Printf("IsBinarySearchTree: %v\n\n", invalid.IsBinarySearchTree())

	// Invalid: 25 is in left subtree of 20 but 25 > 20 (only min/max catches this)
	invalidAncestor := &tree.BinarySearchTree[int]{
		Root: &tree.Node[int]{
			Value: 20,
			Left: &tree.Node[int]{
				Value: 10,
				Right: &tree.Node[int]{Value: 25},
			},
			Right: &tree.Node[int]{Value: 30},
		},
	}
	fmt.Println("Invalid BST (25 in left subtree of 20):")
	invalidAncestor.Print()
	fmt.Printf("IsBinarySearchTree: %v\n", invalidAncestor.IsBinarySearchTree())
}

