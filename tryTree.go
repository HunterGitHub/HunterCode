package main

import (
	"fmt"
)

type E int16

const (
	RED = 0
	BLACK = 1
)

type TreeNode struct {
	E
	color	int
	p		*TreeNode
	left	*TreeNode
	right	*TreeNode
}

type Tree struct {
	root	*TreeNode
}

/*
func (t *Tree) Insert(e E) bool {
	var x = t.root
	var y *TreeNode
	for x != nil {
		y = x
		if e < x.E {
			x = x.left
		} else {
			x = x.right
		}
	}
	n := &TreeNode{e, RED, y, nil, nil}
	if y == nil {
		t.root = n
	} else if e < y.E {
		y.left = n
	} else {
		y.right = n
	}
	// fix up
	if n.p == nil {
		n.color == BLACK
	} else if n.p.color == BLACK {
		return true
	}
	for n.p.color == RED {
		if n.p == n.p.p.left {
			uncle := n.p.p.right
			if uncle.color == RED {
				n.p.color	= BLACK
				uncle.color = BLACK
				n.p.p.color	= RED
				n			= n.p.p
			} else if n == n.p.right {
				n = n.p
				L_ROTATE(n)
			} else if n == n.p.left {
				n.p.color	= BLACK
				n.p.p.color	= RED
				R_ROTATE(n.p.p)
			}
		} else if n.p == n.p.p.right {
			uncle := n.p.p.left
			if uncle.color == RED {
				n.p.color	= BLACK
				uncle.color	= BLACK
				n.p.p.color	= RED
				n			= n.p.p
			} else if n == n.p.left {
				n = n.p
				R_ROTATE(n)
			} else if n == n.p.right {
				n.p.color	= BLACK
				n.p.p.color = RED
				L_ROTATE(n.p.p)
			}
		}
	}
	t.root.color = BLACK

	return true
}
*/

func (t *Tree) remove() {
}

func (t *TreeNode) printT() {
	if t.left != nil {
		fmt.Print("<-")
		t.left.printT()
	}
	fmt.Println(t.E)
	if t.right != nil {
		fmt.Print("->")
		t.right.printT()
	}
}

func main() {
	var t = &Tree{nil}
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.root.printT()
}













