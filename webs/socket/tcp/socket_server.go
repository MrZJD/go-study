package main

import (
	"fmt"
	"log"
	"net"
	"strconv"
	"strings"
	"time"
)

// TCP in Server

// func ListenTCP(network string, laddr *TCPAddr) (*TCPAddr, error)
// func (I *TCPListener) Accept() (Conn, error)

func serveTCP() {
	service := ":7777"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		fmt.Println("Resolve Connection")
		daytime := time.Now().String()
		conn.Write([]byte(daytime))
		conn.Close()
	}
}

func serveInRoutine() {
	service := ":1200"
	tcpAddr, err := net.ResolveTCPAddr("tcp4", service)
	if err != nil {
		log.Fatal(err)
	}

	listener, err := net.ListenTCP("tcp", tcpAddr)
	if err != nil {
		log.Fatal(err)
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}
		// go handleClient(conn)
		go handleLongClient(conn)
	}
}

/* 短连接 */
func handleClient(conn net.Conn) {
	defer conn.Close()

	fmt.Println("[Resolve TCP Conn in routine]")

	daytime := time.Now().String()
	conn.Write([]byte(daytime))
}

/* 长连接 */
func handleLongClient(conn net.Conn) {
	conn.SetReadDeadline(time.Now().Add(10 * time.Second))
	request := make([]byte, 128)
	defer conn.Close()

	fmt.Println("[Resolve TCP Conn in longHandler]")

	for {
		readLen, err := conn.Read(request)

		if err != nil {
			fmt.Println(err)
			break
		}

		fmt.Println(strings.TrimSpace(string(request[:readLen])))

		if readLen == 0 {
			break // end of stream
		} else if strings.TrimSpace(string(request[:readLen])) == "timestamp" {
			daytime := strconv.FormatInt(time.Now().Unix(), 10)
			conn.Write([]byte(daytime))
		} else {
			daytime := time.Now().String()
			conn.Write([]byte(daytime))
		}

		request = make([]byte, 128) // clear content
	}
}

func main() {
	// serveTCP()
	serveInRoutine()
}
