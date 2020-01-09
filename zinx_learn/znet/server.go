package znet

import (
	"fmt"
	"net"

	"github.com/yhkl-dev/Notes/zinx_learn/zface"
)

// 接口实现
type Server struct {
	// server name
	Name string
	// server bind ip
	IPVersion string
	// server listen ip
	IP string
	// listen port
	Port int
}

// Start  server
func (s *Server) Start() {
	fmt.Printf("[Start] Server listener at IP: %s, Port: %d, is starting\n", s.IP, s.Port)

	go func() {

		addr, err := net.ResolveTCPAddr(s.IPVersion, fmt.Sprintf("%s:%d", s.IP, s.Port))
		if err != nil {
			fmt.Println("resolve tcp addr error: ", err)
			return
		}

		listenser, err := net.ListenTCP(s.IPVersion, addr)
		if err != nil {
			fmt.Println("Listen ", s.IPVersion, "err")
			return
		}

		fmt.Println("start Zinx server successful", s.Name, "Listening...")

		for {
			conn, err := listenser.AcceptTCP()
			if err != nil {
				fmt.Println("Accept error", err)
				continue
			}

			go func() {
				for {
					buf := make([]byte, 512)
					cnt, err := conn.Read(buf)
					if err != nil {
						fmt.Println("receive buf err", err)
						continue
					}

					fmt.Printf("receive client buf %s, count %d\n", buf, cnt)
					if _, err := conn.Write(buf[:cnt]); err != nil {
						fmt.Println("write back buf err")
						continue
					}
				}
			}()
		}

	}()
}

// Stop server
func (s *Server) Stop() {
	// todo
}

func (s *Server) Serve() {
	s.Start()

	select {}
}

// NewServer server init func
func NewServer(name string) zface.Zserver {
	s := &Server{
		Name:      name,
		IPVersion: "tcp4",
		IP:        "0.0.0.0",
		Port:      9999,
	}
	return s
}
