package main

import (
	"net"
	"fmt"
	"syscall"
	"time"
	"bytes"
)

func main() {
	fmt.Println(net.InterfaceAddrs())
	l, _ := net.Interfaces()
	for _, v := range l {
		fmt.Println(v)
	}
	fd, err := syscall.Socket(syscall.AF_INET, syscall.SOCK_PACKET, 0x0300)
	if err != nil {
		fmt.Println("socket", err)
	}
	defer syscall.Close(fd)

	if err = syscall.BindToDevice(fd, "eth1"); err != nil {
		fmt.Println("bind", err)
		panic(err)
	}

	for {
		b := make([]byte, 100)
		_, _, _ = syscall.Recvfrom(fd, b, 0)
		if ! IsIcmpPac(b) {
			continue
		}
		//fmt.Printf("%5x\n", b[12:14])
		//continue
		for k, v := range b {
			if k % 4 == 2 {
				fmt.Println()
			}
			fmt.Printf("|%4d: %5d:0x%-5x|", k, v, v)
		}
		fmt.Println("\na packet-----------")
	}
	time.Sleep(1 * time.Second)
	fmt.Println(fd)

}

func IsIpPac(pack []byte) bool {
	return bytes.Equal(pack[12:14], []byte{8, 0})
}

func IsArpPac(pack []byte) bool {
	return bytes.Equal(pack[12:14], []byte{8, 6})
}

func IsIcmpPac(pack []byte) bool {
	//fmt.Println("the 23", pack[23])
	if bytes.Equal(pack[12:14], []byte{8, 0}) && pack[23] == 1 {
		return true
	} else {
		return false
	}
}
