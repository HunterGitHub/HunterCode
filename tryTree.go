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

func (t *Tree) L_Rotate(n *TreeNode) {
	y := n.right
	n.right = y.left		// child switch, x, y independent
	if y.left != nil {
		y.left.p = n		// child chg parent
	}
	y.p = n.p				// middle chg parent
	if n.p == nil {
		t.root = y
	} else if n == n.p.left {
		n.p.left = y			// y switch
	} else if n == n.p.right {
		n.p.right = y			// y switch
	}
	y.left = n
	n.p = y
}

func (t *Tree) R_Rotate(n *TreeNode) {
	y := n.left
	n.left = y.right
	if y.right != nil {
		y.right.p = n
	}
	y.p = n.p
	if n.p == nil {
		t.root = y
	} else if n == n.p.right {
		n.p.right = y
	} else if n == n.p.left {
		n.p.left = y
	}
	y.right = n
	n.p = y
}

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
		n.color = BLACK
		return true
	} else if n.p.color == BLACK {
		return true
	}
	for n.p.color == RED {
		if n.p == n.p.p.left {
			uncle := n.p.p.right
			if uncle != nil && uncle.color == RED {
				n.p.color	= BLACK
				uncle.color = BLACK
				n.p.p.color	= RED
				n			= n.p.p
			} else if n == n.p.right {
				n = n.p
				t.L_Rotate(n)
			} else if n == n.p.left {
				n.p.color	= BLACK
				n.p.p.color	= RED
				t.R_Rotate(n.p.p)
			}
		} else if n.p == n.p.p.right {
			uncle := n.p.p.left
			if uncle != nil && uncle.color == RED {
				n.p.color	= BLACK
				uncle.color	= BLACK
				n.p.p.color	= RED
				n			= n.p.p
			} else if n == n.p.left {
				n = n.p
				t.R_Rotate(n)
			} else if n == n.p.right {
				n.p.color	= BLACK
				n.p.p.color = RED
				t.L_Rotate(n.p.p)
			}
		}
		if n.p == nil {
			break;
		}
	}
	t.root.color = BLACK

	return true
}

func (t *Tree) remove() {
}

func (t *TreeNode) printT() {
	fmt.Println(t.E, t.color)
	if t.left != nil {
		fmt.Print("<-")
		t.left.printT()
	}
	fmt.Println("----")
	if t.right != nil {
		fmt.Print("->")
		t.right.printT()
	}
}

func (t *TreeNode) printGood() {
	var q = make(chan *TreeNode, 50)
	q <- t
	for {
		select {
		case v := <-q:
			if v != nil {
				fmt.Println(v.E, v.color, v.p)
				q <- v.left
				q <- v.right
			}
		default:
			fmt.Println("end--")
			goto end
		}
	}
	end:
}

func main() {
	var t = &Tree{nil}
	t.Insert(1)
	t.Insert(2)
	t.Insert(3)
	t.Insert(4)
	t.Insert(5)
	t.Insert(6)
	t.Insert(7)
	t.Insert(8)
	t.Insert(9)
	t.root.printT()
	t.root.printGood()
}













