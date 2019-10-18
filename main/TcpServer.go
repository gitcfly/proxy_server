package main

import (
	"fmt"
	"net"
)

//go-tcpsock/server.go
func handleConn(c net.Conn) {
	defer func() {
		c.Close()
		fmt.Printf("Close a connetion src address: %v \n\n", c.RemoteAddr().String())
	}()
	fmt.Printf("Recive a new connetion src address: %v \n", c.RemoteAddr().String())
	var data = make([]byte, 65500)
	len, err := c.Read(data)
	if err != nil {
		fmt.Println(err)
	}
	if len > 0 {
		fmt.Println(string(data))
	}
	_, err = c.Write([]byte("hello client,i am server!"))
	if err != nil {
		fmt.Println(err)
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
