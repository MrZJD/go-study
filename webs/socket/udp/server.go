package main

import (
	"fmt"
	"log"
	"net"
	"time"
)

// -> TCP与UDP在语言层面上基本一致
// 唯一的不同: UDP缺少了对客户端连接请求的Accept函数

// func ResolveUDPAddr(net, addr string) (*UDPAddr, os.Error)
// func DialUDP(net string, laddr, raddr *UDPAddr) (c *UDPConn, err os.Error)
// func ListenUDP(net string, laddr *UDPAddr) (c *UDPConn, err os.Error)
// func (c *UDPConn) ReadFromUDP(b []byte) (n int, addr *UDPAddr, err os.Error)
// func (c *UDPConn) WriteToUDP(b []byte, addr *UDPAddr) (n int, err os.Error)

func main() {
	service := ":1300"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	for {
		handleClient(conn)
	}
}

func handleClient(conn *net.UDPConn) {
	var buf [512]byte

	_, addr, err := conn.ReadFromUDP(buf[0:])
	if err != nil {
		return
	}

	fmt.Println("[Handler Request from UDP]")

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
