package main

import (
	"fmt"

	"github.com/xxmyjk/go-alg/pkg/matrix"
	"github.com/xxmyjk/go-alg/pkg/tree"
)

func main() {
	// CircleLikePrint matrix
	el := []int{
		1, 2, 3, 4, 5, 6, 7,
		8, 9, 10, 11, 12, 13, 14,
		15, 16, 17, 18, 19, 20, 21,
		50, 23, 24, 25, 26, 27, 28,
	}

	m := matrix.New(4, 7, el...)
	m.CircleLikePrint()

	// Reverse tree
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

	root := &tree.Node{Value: 1, Left: nil, Right: nil}
	root.Left = &tree.Node{Value: 2, Left: nil, Right: nil}
	root.Left.Left = &tree.Node{Value: 4, Left: nil, Right: nil}
	root.Left.Right = &tree.Node{Value: 5, Left: nil, Right: nil}

	root.Right = &tree.Node{Value: 3, Left: nil, Right: nil}
	root.Right.Left = &tree.Node{Value: 6, Left: nil, Right: nil}
	root.Right.Right = &tree.Node{Value: 7, Left: nil, Right: nil}

	tree.Print(root)
	tree.Reverse(root)
	fmt.Println()
	tree.Print(root)
}
