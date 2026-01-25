package chapter2

import (
	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared/stack"
)

// RunChapter2Exercise6
// Palindrome: Implement a function to check if a linked list is a palindrome.
// The list is represented as a singly linked list where each node contains a single value.
//
// Example:
// Input:  1 -> 2 -> 3 -> 2 -> 1
// Output: true
func RunChapter2Exercise6() {
	// TODO: Build example inputs and call isPalindrome.
	exampleSlice := []int{1, 2, 3, 2, 1}
	list := shared.NewLinkedListFromSlice(exampleSlice)
	list.Print()
	result := isPalindrome(list.Head)
	println("Is palindrome?", result)
}

// Time: O(N/2) + O(N/2) -> O(N)
// Space: O(N/2) -> O(N) for stack created to revert seconf half valeus
func isPalindrome(head *shared.Node[int]) bool {
	if head == nil {
		return false
	}

	slow := head
	fast := head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}

	middleNode := slow
	firstHalfPointer := head
	secondHalfPointer := middleNode

	secondHalfStack := stack.Stack[int]{}
	for secondHalfPointer != nil {
		secondHalfStack.Push(secondHalfPointer.Value)
		secondHalfPointer = secondHalfPointer.Next
	}

	for {
		poppedV, _ := secondHalfStack.Pop()
		if firstHalfPointer.Value != poppedV {
			return false
		}

		if firstHalfPointer == middleNode || secondHalfStack.IsEmpty() {
			break
		}

		firstHalfPointer = firstHalfPointer.Next
	}

	return true
}
