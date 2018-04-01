package module

import (
	"net"
	"fmt"
)

const (
	// Login = 1
	//Register = 2
	//Send = 3
	//UserList = 4
)

func LoginCmd(conn net.Conn)  {
	buf := make([]byte, 512)
	conn.Write([]byte("please input you id"))
	n, _ := conn.Read(buf)
	uid := buf[:n]
	conn.Write([]byte("please input you password"))
	n, _ = conn.Read(buf)
	passwd := buf[:n]
	if string(uid) == string(passwd) {
		fmt.Println("login success")
		conn.Write([]byte("OK"))

	}else {
		conn.Write([]byte("login filed: password or username error"))
	}

}

