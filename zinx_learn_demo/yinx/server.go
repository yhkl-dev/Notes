package main

import "github.com/yhkl-dev/Notes/zinx_learn/znet"

func main() {
	s := znet.NewServer("yhkl-v1.0")
	s.Serve()
}
