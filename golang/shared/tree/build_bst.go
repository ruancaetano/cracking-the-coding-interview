package tree

func BuildSearchTreeFromSortedArray[T any](input []T) *BinarySearchTree[T] {
	if len(input) == 0 {
		return nil
	}

	root := buildSearchTreeFromSortedArray(input)
	return &BinarySearchTree[T]{
		Root: root,
	}
}

func buildSearchTreeFromSortedArray[T any](input []T) *Node[T] {
	if len(input) == 0 {
		return nil
	}

	if len(input) == 1 {
		return &Node[T]{
			Value: input[0],
		}
	}

	middle := len(input) / 2
	left := input[:middle]
	right := input[middle+1:]

	root := &Node[T]{
		Value: input[middle],
		Left:  buildSearchTreeFromSortedArray(left),
		Right: buildSearchTreeFromSortedArray(right),
	}

	return root
}
