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
	b interface{
		haha() int
	}
}

func (my mystruct) foo() int {
	fmt.Println("fdfd....fd")
	return 11
}

type myinter interface {
	a()
}

func main() {
	//var path = "/home/admin/aitao"
	//lib.DialGoogle()
	mywords := []string{"hi","im shit","oh shit"}
	//myint := []int{1,2,3}
	fmt.Println(mywords)
	var a = mystruct{1,2, func(){ mystruct.foo(mystruct{}) }}
	//fmt.Println(a.a())
	mystruct.foo(mystruct{})
	//var mi myinter = mystruct{}
	//mi.a()
	//myinter.foo()

	a.a()
	pr(mywords)
	
}
