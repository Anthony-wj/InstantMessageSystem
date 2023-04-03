/*
	server类型 Ip 和 Port

			创建一个Server对象
	方法    启动Server服务
	  		处理链接业务

*/

package main

import (
	"fmt"
	"net"
)

type Server struct {
	Ip   string
	Port int
}

// 创建一个server的接口
func NewServer(ip string, port int) *Server {
	server := &Server{
		Ip:   ip,
		Port: port,
	}
	return server
}

func (s *Server) handler(conn net.Conn) {
	// 当前连接的业务
	fmt.Println("连接建立成功")
}

func (s *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%d", s.Ip, s.Port))
	if err != nil {
		fmt.Printf("net.Listen err: %v\n", err)
		return
	}

	// close listen socket
	defer listener.Close()
	for {
		// accept
		conn, err := listener.Accept()
		if err != nil {
			fmt.Printf("listener accept err: %v\n", err)
			continue
		}
		// do handler
		go s.handler((conn))
	}

}
