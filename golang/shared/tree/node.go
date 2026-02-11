package tree

type Node[T comparable] struct {
	Value T
	Left  *Node[T]
	Right *Node[T]
	Depth int
}
