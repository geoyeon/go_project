package main

import (
	"fmt"
	"net"
)

//데이터를 msgPack 추후적용
func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		fmt.Println(string(data[:n]))

		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

func main() {
	ln, err := net.Listen("tcp", ":8000")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ln.Close()

	fmt.Println("Start TCP Server")

	for {
		conn, err := ln.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}
		defer conn.Close()

		go requestHandler(conn)
	}
}
