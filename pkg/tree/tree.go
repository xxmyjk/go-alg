package tree

import "fmt"

/**
1
2 | 3
4  5 | 6 7
**/

/**
==>

1
3 | 2
7 6 | 5 4
**/

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func Reverse(n *Node) *Node {
	if n == nil {
		return nil
	}

	n.Left, n.Right = n.Right, n.Left
	Reverse(n.Left)
	Reverse(n.Right)

	return n
}

func Print(n *Node) {
	if n == nil {
		return
	}
	fmt.Print(n.Value, "\t")
	Print(n.Left)
	Print(n.Right)
}
