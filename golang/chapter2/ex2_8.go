package chapter2

import "github.com/ruancaetano/cracking-the-coding-interview/golang/shared"

// RunChapter2Exercise8
// Loop Detection: Given a circular linked list, implement an algorithm that returns
// the node at the beginning of the loop.
//
// Definition:
// A circular linked list is a (corrupt) linked list in which a node's next pointer
// points to an earlier node, so as to make a loop in the linked list.
//
// Example:
// Input:  A -> B -> C -> D -> E -> C (the same C as earlier in the list)
// Output: C
//
// Use this function to build your own circular list examples and test your solution.
func RunChapter2Exercise8() {
	// TODO: Build example circular linked lists and call detectLoopStart.
}

// detectLoopStart should return the node at the beginning of the loop,
// or nil if there is no loop in the list.
// Implement the logic for this function as part of your solution.
func detectLoopStart(head *shared.Node[int]) *shared.Node[int] {
	// TODO: Implement loop detection for the linked list.
	panic("implement me")
}
