package chapter2

import (
	"fmt"

	"github.com/ruancaetano/cracking-the-coding-interview/golang/shared"
)

// RunChapter2Exercise7
// Intersection: Given two (singly) linked lists, determine if the two lists intersect.
// Return the intersecting node. Note that intersection is defined based on reference,
// not value. That is, if the kth node of the first linked list is the exact same node
// (by reference) as the jth node of the second linked list, then they are intersecting.
//
// Example:
// List A: 3 -> 7 -> 8 -> 10
// List B:     99 -> 1 ↘
//
//	8 -> 10
//
// Use this function to construct example lists and test your implementation.
func RunChapter2Exercise7() {
	// TODO: Build example linked lists and call findIntersection.
	aValues := []int{3, 7, 8, 10}
	bValues := []int{99, 1, 8, 10}
	listA := shared.NewLinkedListFromSlice(aValues)
	listB := shared.NewLinkedListFromSlice(bValues)

	refNode := listA.NodeAt(2)
	listB.SetNextAt(2, refNode)

	fmt.Print("List A: ")
	listA.Print()
	fmt.Print("List B: ")
	listB.Print()

	intersecitonNode1 := findIntersection(listA.Head, listB.Head)
	inteserctionNode2 := optimizedFindIntersection(listA.Head, listB.Head)

	fmt.Print("Result: ")
	if intersecitonNode1 != nil {
		intersecitonNode1.Print()
	} else {
		fmt.Println("No intersection")
	}

	fmt.Print("Result with optmized func: ")
	if inteserctionNode2 != nil {
		inteserctionNode2.Print()
	} else {
		fmt.Println("No intersection")
	}

}

// Time: O(M * N)
// Space: O(1)
func findIntersection(a, b *shared.Node[int]) *shared.Node[int] {
	refA := a
	refB := b

	for refA != nil {
		for refB != nil {
			if refB == refA {
				return refB
			}

			refB = refB.Next
		}

		refA = refA.Next
		refB = b
	}

	return nil
}

// Time: O(M + N)
// Space: O(1)
func optimizedFindIntersection(a, b *shared.Node[int]) *shared.Node[int] {
	if a == nil || b == nil {
		return nil
	}

	pa := a
	pb := b

	countA := calcLen(a)
	countB := calcLen(b)

	if countA > countB {
		diff := countA - countB
		for diff > 0 {
			pa = pa.Next
			diff--
		}
	} else {
		diff := countB - countA
		for diff > 0 {
			pb = pb.Next
			diff--
		}
	}

	// pa and pb have the same size from here
	for pa != nil && pb != nil {
		if pa == pb {
			return pa
		}
		pa = pa.Next
		pb = pb.Next
	}

	return nil
}

func calcLen(node *shared.Node[int]) int {
	count := 0

	for node != nil {
		count += 1
		node = node.Next
	}

	return count
}
