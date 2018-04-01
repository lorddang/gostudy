package main

import (
	"fmt"
	"encoding/binary"
	"bytes"
)

func main() {
	conf := NewConfig()
	conf.DBName = "test"
	conf.User = "root"
	conf.Addr = "localhost:3306"
	conf.Passwd = "123456"

	mc := &mysqlConn{
		buf:              newBuffer(nil),
		maxAllowedPacket: maxPacketSize,
		cfg: &Config{
			InterpolateParams: true,
		},
	}
	mc.readInitPacket()
}

func readInitPacket(p []byte) (error) {
	data := p
	if data[0] == iERR {
		fmt.Println("data[0] is iError")
		return nil
	}
	if data[0] < minProtocolVersion {
		fmt.Println("data[0] < minProtocolVersion")
		return nil
	}
	pos := 1 + bytes.IndexByte(data[1:], 0x00) + 1 + 4
	fmt.Println(pos)
	fmt.Println(bytes.IndexByte(data[1:], 0x00))
	cipher := data[pos:pos + 8]

	pos += 8 + 1

	fmt.Printf("chiper is data[%d, %d] %x", pos, pos + 8, cipher)
	flags := clientFlag(binary.LittleEndian.Uint16(data[pos : pos+2]))
	if flags&clientProtocol41 == 0 {
		return nil
	}
	if flags&clientSSL == 0 {
		return nil
	}
	pos += 2

	if len(data) > pos {
		// character set [1 byte]
		// status flags [2 bytes]
		// capability flags (upper 2 bytes) [2 bytes]
		// length of auth-plugin-data [1 byte]
		// reserved (all [00]) [10 bytes]
		pos += 1 + 2 + 2 + 1 + 10

		// second part of the password cipher [mininum 13 bytes],
		// where len=MAX(13, length of auth-plugin-data - 8)
		//
		// The web documentation is ambiguous about the length. However,
		// according to mysql-5.7/sql/auth/sql_authentication.cc line 538,
		// the 13th byte is "\0 byte, terminating the second part of
		// a scramble". So the second part of the password cipher is
		// a NULL terminated string that's at least 13 bytes with the
		// last byte being NULL.
		//
		// The official Python library uses the fixed length 12
		// which seems to work but technically could have a hidden bug.
		cipher = append(cipher, data[pos:pos+12]...)

		// TODO: Verify string termination
		// EOF if version (>= 5.5.7 and < 5.5.10) or (>= 5.6.0 and < 5.6.2)
		// \NUL otherwise
		//
		//if data[len(data)-1] == 0 {
		//	return
		//}
		//return ErrMalformPkt

		// make a memory safe copy of the cipher slice
		var b [20]byte
		copy(b[:], cipher)
		fmt.Printf("挑战随机数 %x \n", cipher)
		return nil
	}

	// make a memory safe copy of the cipher slice
	var b [8]byte
	copy(b[:], cipher)
	fmt.Printf("挑战随机数 %x \n", cipher)
	return nil

}
