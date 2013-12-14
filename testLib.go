package main

import (
	"net"
	"fmt"
)

func main() {
	ip := "172.23.100.10"
	ip = "172.21.98.215"
	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupAddr("172.24.108.35"))
	fmt.Println(net.LookupAddr(ip))
}
