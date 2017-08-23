package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"net"
)

type Item struct {
	Msg string
	Num int
}

//데이터를 msgPack 적용
func requestHandler(c net.Conn) {
	data := make([]byte, 4096)

	for {
		n, err := c.Read(data)
		if err != nil {
			fmt.Println(err)
			return
		}

		var item Item
		err = msgpack.Unmarshal(data[:n], &item)
		if err != nil {
			panic(err)
		}

		fmt.Println(item)

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
