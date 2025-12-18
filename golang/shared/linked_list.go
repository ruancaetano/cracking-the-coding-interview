package shared

type LinkedList[T comparable] struct {
	Head *Node[T]
}

type Node[T comparable] struct {
	Value T
	Next  *Node[T]
}

// PrintNode prints the value of the given node and all subsequent nodes in the list.
func (node *Node[T]) Print() {
	current := node
	for current != nil {
		print(current.Value)
		if current.Next != nil {
			print(" -> ")
		}
		current = current.Next
	}
	println()
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

func (l *LinkedList[T]) NodeAt(index int) *Node[T] {
	current := l.Head
	pos := 0
	for current != nil {
		if pos == index {
			return current
		}
		current = current.Next
		pos++
	}
	return nil
}

func (l *LinkedList[T]) SetNextAt(index int, node *Node[T]) {
	if index == 0 {
		l.Head = node
		return
	}
	prev := l.Head
	pos := 0
	for prev != nil && pos < index-1 {
		prev = prev.Next
		pos++
	}
	if prev != nil && prev.Next != nil {
		prev.Next = node
	}
}
