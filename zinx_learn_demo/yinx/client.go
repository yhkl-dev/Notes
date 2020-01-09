package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	fmt.Println("client start")
	time.Sleep(time.Second * 1)

	conn, err := net.Dial("tcp", "127.0.0.1:9999")
	if err != nil {
		fmt.Println("client start error", err)
		return
	}

	for {
		_, err := conn.Write([]byte("hello yhkl 1,0"))
		if err != nil {
			fmt.Println("write conn error", err)
			return
		}

		buf := make([]byte, 512)
		count, err := conn.Read(buf)
		if err != nil {
			fmt.Println("read buf error", err)
			return
		}

		fmt.Printf("server call back %s, count=%d\n", buf, count)
		time.Sleep(time.Second * 1)
	}

}
