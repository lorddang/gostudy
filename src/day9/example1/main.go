package main

import (
	"fmt"
	"net"
	"log"
)

func main()  {
	fmt.Println("start server")
	listener, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatalf("Listen 8000 error %v", err)
		return
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			log.Fatalf("accept 8000 error %v", err)
			continue
		}
		go processAccept(conn)

	}
}

func processAccept(conn net.Conn)  {
	defer conn.Close()
	for {
		buf := make([]byte, 5120)
		n, err := conn.Read(buf)
		if err != nil{

			log.Fatalf("read error %v", err)
			break
		}
		fmt.Printf("> %s \n", string(buf[:n]))
	}


}


