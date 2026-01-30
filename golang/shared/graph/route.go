package graph

import (
	"fmt"
)

func PrintRoute(label string, path []*Node[int]) {
	if len(path) == 0 {
		fmt.Printf("  %s: no route\n", label)
		return
	}
	vals := make([]int, len(path))
	for i, n := range path {
		vals[i] = n.GetValue()
	}
	fmt.Printf("  %s: %v\n", label, vals)
}
