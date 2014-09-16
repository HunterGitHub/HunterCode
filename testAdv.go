package main

import (
	"time"
	"fmt"
)

func main() {
	var _ = make(chan int)
	select {
		case <-time.After(5 * time.Second):
			fmt.Println("timed out")
	}
}

func handle(a int) {
	fmt.Println("---", a)
}
