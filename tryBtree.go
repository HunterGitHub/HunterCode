package main

import (
	"fmt"
)

const MAX = 5
const MIN = 3

type BNode struct {
	num			int16
	ele			[MAX-1]string
	bptr		[MAX]*BNode
}

func NewBtree() *BNode {
	return &BNode{num:0}
}

func (bn *BNode) insert() error {
	return nil
}

func main() {
	bn := NewBtree()
	fmt.Println(bn)
}
