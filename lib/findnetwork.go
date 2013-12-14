package lib

import (
	"net"
	"bufio"
	"fmt"
)

func DialGoogle() {
	conn, err := net.Dial("tcp", "google.com:80")
	if err != nil {
		panic("dial google down!")
	}
	fmt.Fprintf(conn, "GET / HTTP/1.1\r\n\r\n")
	status, err := bufio.NewReader(conn).ReadString('\n')
	fmt.Println(status)
}
