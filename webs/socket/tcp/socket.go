package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

// IP

func useIP() {
	if len(os.Args) != 2 {
		fmt.Fprintf(os.Stdout, "Usage: %s ip-addr\n ", os.Args[0])
		os.Exit(1)
	}

	name := os.Args[1]
	addr := net.ParseIP(name)
	if addr == nil {
		fmt.Println("Invalid Address")
	} else {
		fmt.Println("The Address is ", addr.String())
	}
	os.Exit(0)
}

// TCP
// type TCPAddr struct {
// 	IP   IP
// 	Port int
// 	Zone string // ipv6
// }
// func ResolveTCPAddr(net, addr string) (*TCPAddr, os.Error) {} // 获取TCPAddr

// TCPConn 数据交互
// func (c *TCPConn) Write(b []byte) (int, error)
// func (c *TCPConn) Read(b []byte) (int, error)

// get TCPConn
// func DialTCP(net string, laddr, raddr *TCPAddr) (*TCPConn, error)

// control TCP
// func DialTimeout(net, addr string, timeout time.Duration) (Conn, error)

// func (c *TCPConn) SetReadDeadline(t time.Time) error
// func (c *TCPConn) SetWriteDeadline(t time.Time) error

// func (c *TCPConn) SetKeepAlive(keepalive bool) os.Error

// In Client
func RequestTCP() {
	tcpAddr, err := net.ResolveTCPAddr("tcp4", ":1200")
	if err != nil {
		fmt.Println("[ResolveTCPAddr]")
		log.Fatal(err)
	}

	conn, err := net.DialTCP("tcp", nil, tcpAddr)
	if err != nil {
		fmt.Println("[DialTCP]")
		log.Fatal(err)
	}

	_, err = conn.Write([]byte("HEAD /HTTP/1.0\r\n\r\n"))
	if err != nil {
		fmt.Println("[conn Write]")
		log.Fatal(err)
	}

	// time.Sleep(1 * time.Second)
	// _, err = conn.Write([]byte("timestamp"))
	// if err != nil {
	// 	fmt.Println("[conn Write2]")
	// 	log.Fatal(err)
	// }
	request := make([]byte, 128)
	for {
		lenS, err := conn.Read(request) // 这里没有数据应该是阻塞了

		if err != nil {
			fmt.Println(err)
			break
		}

		if lenS == 0 {
			break // end of stream
		}

		fmt.Println(strings.TrimSpace(string(request[:lenS])))

		request = make([]byte, 128) // clear content
	}

	// result, err := ioutil.ReadAll(conn) // 等待数据传输结束
	// if err != nil {
	// 	fmt.Println("[ioutil.ReadAll]")
	// 	log.Fatal(err)
	// }

	// fmt.Println(string(result))
}

func main() {
	// useIP()
	RequestTCP()
}
