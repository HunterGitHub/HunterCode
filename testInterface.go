package main

import (
	"fmt"
)

type abc struct {
	A int
}

func main() {
	var a interface{} = abc{1}
	var b abc = abc{1}
	fmt.Println(b.A)
}
