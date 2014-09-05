package main

import (
	"fmt"
)

const BUCKET_SIZE = 5
const BUCKET_MAX  = 10

type DiyMap struct {
	table	[]*Node
}

type Node struct {
	k		string
	v		string
	next	*Node
}

func NewDiyMap() (r DiyMap) {
	r = DiyMap{}
	r.table = make([]*Node, BUCKET_SIZE, BUCKET_MAX)
	return r
}

func (dm *DiyMap) k(k string) (string, bool) {
	hash := len(k) % BUCKET_SIZE
	n := dm.table[hash]
	for ; n != nil; n = n.next {
		if n.k == k {
			return n.v, true
		}
	}
//	fmt.Println("none")
	return "", false
}

func (dm *DiyMap) kv(k string, v string) error {
	hash := len(k) % BUCKET_SIZE
	if dm.table[hash] == nil {
		dm.table[hash] = &Node{k, v, nil}
		return nil
	}
	n := dm.table[hash]
	for n != nil {
		if n.k == k && n.v == v {
			fmt.Println("already set kv")
			return nil
		}
		if n.next == nil {
			n.next = &Node{k, v, nil};
			fmt.Println("set kv: ", n.next)
			return nil
		} else {
			n = n.next
		}
	}
	return nil
}

func main() {
	r := NewDiyMap()
	r.kv("fdfd", "yes")
	r.kv("shit", "nono")
	fmt.Println(r.k("fdfd"))
	fmt.Println(r.k("shit"))
}
