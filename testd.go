package main

import (
)

type aa struct {
}

func (q aa) a(a int) {
}

func main() {
	aa.a(aa{})
}
