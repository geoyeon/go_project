package main

import (
	"fmt"
	"github.com/vmihailenco/msgpack"
	"net"
	"time"
)

type Item struct {
	Msg string
	Num int
}

//데이터 통신을 msgPack 적용
func main() {
	client, err := net.Dial("tcp", "127.0.0.1:8000")

	if err != nil {
		fmt.Println(err)
		return
	}
	defer client.Close()

	go func(c net.Conn) {
		data := make([]byte, 4096)

		for {
			n, err := c.Read(data)
			if err != nil {
				fmt.Println(err)
				return
			}

			fmt.Println(data[:n])

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

			time.Sleep(1 * time.Second)
		}
	}(client)

	go func(c net.Conn) {
		i := 0
		for {
			item, err := msgpack.Marshal(&Item{Msg: "Hello : ", Num: i})
			if err != nil {
				panic(err)
			}

			c.Write(item)
			if err != nil {
				fmt.Println(err)
				return
			}

			i++
			time.Sleep(1 * time.Second)
		}
	}(client)

	fmt.Scanln()
}
