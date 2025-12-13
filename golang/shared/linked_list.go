package shared

type LinkedList[T comparable] struct {
	Head *Node[T]
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

func NewLinkedListFromSlice[T comparable](slice []T) LinkedList[T] {
	list := LinkedList[T]{}
	if len(slice) == 0 {
		return list
	}

	list.Head = &Node[T]{Value: slice[0]}
	current := list.Head
	for _, v := range slice[1:] {
		current.Next = &Node[T]{Value: v}
		current = current.Next
	}
	return list
}

func (l *LinkedList[T]) Print() {
	current := l.Head
	for current != nil {
		print(current.Value)
		if current.Next != nil {
			print(" -> ")
		}
		current = current.Next
	}
	println()
}
