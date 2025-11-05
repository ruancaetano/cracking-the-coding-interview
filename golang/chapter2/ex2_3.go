package chapter2

// RunChapter2Exercise3
// Input: A singly linked list and a node in the middle (e.g., 1 -> 2 -> 3 -> 4 -> 5, node=3)
// The task: Remove a node from the middle of a singly linked list, given only access to that node.
// (You do NOT have access to the head of the list.)

// Example input: 1 -> 2 -> 3 -> 4 -> 5 (target node: 3)
// Expected output after removal: 1 -> 2 -> 4 -> 5

import (
	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

func RunChapter2Exercise3() {
	input := util.NewLinkedListFromSlice([]int{1, 2, 3, 4, 5})
	input.Print()

	// ref to node wit value 3
	middle := input.Head.Next.Next // this points to node with value 3

	removeNode(middle) // This would be the function to solve the exercise

	input.Print() // Expected: 1 -> 2 -> 4 -> 5
}

// Time: O(1)
// Space: O(1)
func removeNode(node *util.Node[int]) {
	node.Value = node.Next.Value
	node.Next = node.Next.Next
}
