package main

import (
	"fmt"
)

type LinkList interface {
	insert(ele string) bool
	remove(ele string) bool
}

type Node struct {
	ele		string
	next	*Node
}

func NewLinkList() *Node {
	return &Node{"head", nil}
}

func (head *Node) insert(ele string) bool {
	n := &Node{ele, nil}
	n.next = head.next
	head.next = n
	return true
}

func (head *Node) remove(ele string) bool {
	n := head
	for n.next != nil {
		if n.next.ele == ele {
			//remove
			n.next = n.next.next
			return true
		}
		n = n.next
	}
	return false
}

func (head *Node) find(ele string) *Node{
	n := head.next
	for n != nil {
		if n.ele == ele {
			return n
		}
		n = n.next
	}
	return nil
}

func (head *Node) printL() {
	fmt.Print(head.ele, ": ")
	for n := head.next; n != nil; n = n.next {
		fmt.Print(n.ele, "=> ")
	}
	fmt.Println()
}

type DulNode struct {
	ele		string
	pre		*DulNode
	next	*DulNode
}

func NewDulNode(ele string) *DulNode {
	head := &DulNode{ele, nil, nil}
	return head
}

func (head *DulNode) insert(ele string) bool {
	n := &DulNode{ele, nil, nil}
	if head.next == nil {
		head.next, head.pre = n, n
		n.next = n
		n.pre = n
		return true
	} else {
		n.next = head.next
		n.pre  = head.next.pre
		head.next.pre.next = n
		head.next.pre = n
		return true
	}
	return false
}

func (head *DulNode) remove(ele string) bool {
	return false
}

func (head *DulNode) printDul() {
	fmt.Print(head.ele, ": ")
	end := head.next.pre
	for n := head.next; n != end; n = n.next {
		fmt.Print(n.ele, " <- -> ")
	}
	fmt.Print(end.ele, " <- -> ")
	fmt.Println(end.next.ele, "(Round back)")
}

func main() {
	head := NewLinkList()
	head.insert("shit")
	head.insert("bitch")
	head.insert("fuck")
	head.insert("asshole")
	head.remove("fuck")
	head.printL()
	fmt.Println(head.find("bitch"))

	dulN := NewDulNode("bitch")
	dulN.insert("seven")
	dulN.insert("six")
	dulN.insert("five")
	dulN.insert("four")
	dulN.insert("three")
	dulN.insert("two")
	dulN.printDul()
}
