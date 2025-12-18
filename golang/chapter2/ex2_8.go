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
	// Example: A -> B -> C -> D -> E -> C
	list := shared.NewLinkedListFromSlice([]rune{'A', 'B', 'C', 'D', 'E'})

	// Make the list circular: last node (E/5) points to C (3)
	nodeC := list.NodeAt(2)    // C is at index 2
	lastNode := list.NodeAt(4) // E is at index 4
	if lastNode != nil && nodeC != nil {
		lastNode.Next = nodeC
	}

	print("Result list with loop: ")
	loopStart := detectLoopStart(list.Head)
	if loopStart != nil {
		println("Loop starts at node with value:", string(loopStart.Value))
	} else {
		println("No loop detected")
	}

	// Example without loop
	print("Result list withlout loop: ")
	straightList := shared.NewLinkedListFromSlice([]rune{'F', 'G', 'H', 'I'})
	loopStart2 := detectLoopStart(straightList.Head)
	if loopStart2 != nil {
		println("Loop starts at node with value:", string(loopStart2.Value))
	} else {
		println("No loop detected")
	}
}

// Time: O(N)
// Space: O(N)
func detectLoopStart(head *shared.Node[rune]) *shared.Node[rune] {
	visitedSet := make(map[*shared.Node[rune]]struct{})

	ref := head
	for ref != nil {
		if _, found := visitedSet[ref]; found {
			return ref
		}
		visitedSet[ref] = struct{}{}
		ref = ref.Next
	}

	return nil
}
