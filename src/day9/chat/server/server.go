package main

import (
	"net"
	"fmt"
	"log"
	"io"
	"day9/chat/module"
)

type Server struct {
	host string
	port string
	lsnr net.Listener

}

func NewServer(h, p string) (server *Server) {
	server = &Server{
		host:h,
		port:p,
	}
	addr := fmt.Sprintf("%s:%s", server.host,server.port)
	lsnr, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal("create lsnr error", err)
		return
	}
	server.lsnr = lsnr
	return
}

func (s *Server)run() {
	for {
		conn, err := s.lsnr.Accept()
		if err != nil {
			log.Fatal("accept conn error", err)
			return
		}
		s.processConn(conn)
	}

}

func (s *Server)processConn(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 512)
		n, err := conn.Read(buf)
		if err != nil{
			if err == io.EOF{
				 fmt.Println("close by remote")
				 return
			}
			log.Fatalf("read error %v", err)
			return
		}
		if string(buf[:n]) == "login"{
			s.processData(conn)
		}
		// conn.Write(buf[:n])
	}

}

func (s *Server)processData(conn net.Conn){

	module.LoginCmd(conn)

	return
}



