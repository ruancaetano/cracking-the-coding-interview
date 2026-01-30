package tree

type Node[T any] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
}
