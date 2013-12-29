package main

import (
	"net"
	"fmt"
	//"os"
)

func main() {
	ip := "172.23.100.10"
	ip = "172.21.98.215"
	host := "aitao098215.pre.cm3"
	host = "tmallbuy010096013101.cm3"
	host = "tmallbuy010194172045.cm10"
	fmt.Println(net.LookupAddr("127.0.0.1"))
	fmt.Println(net.LookupAddr("172.24.108.35"))
	fmt.Println(net.LookupAddr(ip))
	fmt.Println(net.LookupHost(host))
	//file, _ := os.Open("/home/admin/aitao/bin/jbossctl")
	data := make([]byte, 1000) 
	//file.Read(data)
	fmt.Println(string(data))
	for _, v := range data {
		if v == '\n' {
			fmt.Print("a newline")
		}
	}
	fmt.Println(net.InterfaceAddrs())
	fmt.Println(net.LookupMX("aitao097120.pre.cm3"))
}
