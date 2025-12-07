package chapter2

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/util"
)

// RunChapter2Exercise5
// Sum Lists: You have two numbers represented by a linked list, where each node contains a single digit.
// The digits are stored in reverse order, such that the 1's digit is at the head of the list.
// Write a function that adds the two numbers and returns the sum as a linked list.
// Example:
// Input: (7 -> 1 -> 6) + (5 -> 9 -> 2)  // 617 + 295
// Output: 2 -> 1 -> 9                   // 912
func RunChapter2Exercise5() {
	// Build example lists using util structures
	firstNumber := util.NewLinkedListFromSlice([]int{7, 1, 6})  // 617
	secondNumber := util.NewLinkedListFromSlice([]int{5, 9, 2}) // 295

	fmt.Print("First number:  ")
	firstNumber.Print()
	fmt.Print("Second number: ")
	secondNumber.Print()

	// Call the function to sum the lists (to be implemented)
	resultAHead := sumLists(firstNumber.Head, secondNumber.Head)
	resultBHead := sumListsRecursivily(firstNumber.Head, secondNumber.Head)

	resultA := util.LinkedList[int]{Head: resultAHead}
	resultB := util.LinkedList[int]{Head: resultBHead}

	fmt.Print("Result interactive:        ")
	resultA.Print()

	fmt.Print("Result recursive:        ")
	resultB.Print()
}

// sumLists
// Setup for the solution:
// Implement this function to add two numbers represented by linked lists (digits in reverse order).
// The function should return the head of a new linked list representing the sum.
// Time: O(N) where N is the length of the bigger number
// Space: O(N)
func sumLists(a, b *util.Node[int]) *util.Node[int] {
	var carry int

	var resultHead *util.Node[int]
	resultNode := resultHead

	i := a
	j := b
	for i != nil || j != nil {
		iValue := 0
		if i != nil {
			iValue = i.Value
		}

		jValue := 0
		if j != nil {
			jValue = j.Value
		}

		value := iValue + jValue + carry

		if value >= 10 {
			carry = value / 10
			value = value % 10
		} else {
			carry = 0
		}

		newNode := util.Node[int]{
			Value: value,
			Next:  nil,
		}

		if resultNode == nil {
			resultNode = &newNode
			resultHead = resultNode
		} else {
			resultNode.Next = &newNode
			resultNode = resultNode.Next
		}

		if i != nil {
			i = i.Next
		}
		if j != nil {
			j = j.Next
		}
	}

	if carry > 0 {
		newNode := util.Node[int]{
			Value: carry,
			Next:  nil,
		}
		resultNode.Next = &newNode
	}

	return resultHead
}

func sumListsRecursivily(a, b *util.Node[int]) *util.Node[int] {
	return sumListsRecHelper(a, b, 0)
}

// sumListsRecHelper recursively adds digits from two linked lists and a carry,
// returning the head of the resulting list.
func sumListsRecHelper(a, b *util.Node[int], carry int) *util.Node[int] {
	if a == nil && b == nil && carry == 0 {
		return nil
	}

	value := carry
	var nextA, nextB *util.Node[int]

	if a != nil {
		value += a.Value
		nextA = a.Next
	}

	if b != nil {
		value += b.Value
		nextB = b.Next
	}

	resultNode := &util.Node[int]{
		Value: value % 10,
	}

	nextCarry := value / 10
	resultNode.Next = sumListsRecHelper(nextA, nextB, nextCarry)

	return resultNode
}
