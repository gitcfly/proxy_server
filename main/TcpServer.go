package main

import (
	"fmt"
	"net"
)

//go-tcpsock/server.go
func handleConn(c net.Conn) {
	defer c.Close()
	var packt = make([]byte, 65535)
	for {
		_, err := c.Read(packt)
		if err != nil {
			break
		}
		fmt.Println(string(packt))
		_, err = c.Write([]byte("hello client,i am server!"))
		if err != nil {
			break
		}
	}
}

func main() {
	l, err := net.Listen("tcp", ":65081")
	if err != nil {
		fmt.Println("listen error:", err)
		return
	}

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println("accept error:", err)
			break
		}
		// start a new goroutine to handle
		// the new connection.
		go handleConn(c)
	}
}
