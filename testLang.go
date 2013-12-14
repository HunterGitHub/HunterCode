package main

import (
	//"./lib"
	"fmt"
)

func pr(a ...interface{}){
	for i,v := range a {
		fmt.Println(i,v)
	}
}

type mystruct struct {
	fuck int
	Fu int
	a func() //int
	B interface{
		haha() int
	}
}

func (my mystruct) foo() int {
	fmt.Println("fdfd....fd")
	return 11
}

func (my mystruct) haha() int {
	fmt.Println("hhhhaa")
	return 22
}

type myinter interface {
	fuck() int
	a()
}

func main() {
	//var path = "/home/admin/aitao"
	//lib.DialGoogle()
	mywords := []string{"hi","im shit","oh shit"}
	pr(mywords)
	//myint := []int{1,2,3}
	fmt.Println(mywords)

	var a = mystruct{1,2, func(){ mystruct.foo(mystruct{})}, mystruct{} }
	//fmt.Println(a.a())
	mystruct.foo(mystruct{})
	//var mi myinter = mystruct{}
	//mi.a()
	//myinter.foo()
	fmt.Println(a.B)
	var amap = map[string]int{"fdf":1}
	amap = make(map[string]int, 10)
	var bmap *map[string]int
	bmap = new(map[string]int)
	fmt.Println(amap)
	if bmap == nil {
		fmt.Println("bmap nil!")
	} else {
		fmt.Println(*bmap)
	}
	//oh := make(int)

	a.a()
	
}
