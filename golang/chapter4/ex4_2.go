package chapter4

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/tree"
)

// RunChapter4Exercise2
// Minimal Tree: Given a sorted (increasing order) array with unique integer
// elements, write an algorithm to create a binary search tree with minimal
// height.
func RunChapter4Exercise2() {
	fmt.Println("Chapter 4 - Exercise 2: Minimal Tree")

	examples := []struct {
		name  string
		input []int
	}{
		{"10 elements", []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{"single element", []int{42}},
		{"two elements", []int{1, 2}},
		{"three elements", []int{1, 2, 3}},
		{"odd length (7)", []int{10, 20, 30, 40, 50, 60, 70}},
		{"empty", []int{}},
	}

	for _, ex := range examples {
		fmt.Printf("\n--- %s: %v ---\n", ex.name, ex.input)
		t := tree.BuildSearchTreeFromSortedArray(ex.input)
		if t == nil {
			fmt.Println("(empty tree)")
			continue
		}
		t.Print()
	}
}
