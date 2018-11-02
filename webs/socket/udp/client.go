package main

import (
	"fmt"
	"log"
	"net"
)

func main() {
	service := ":1300"
	udpAddr, err := net.ResolveUDPAddr("udp4", service)
	if err != nil {
		log.Fatal(err)
		return
	}

	conn, err := net.DialUDP("udp", nil, udpAddr)
	if err != nil {
		log.Fatal(err)
		return
	}

	_, err = conn.Write([]byte("anything"))
	if err != nil {
		log.Fatal(err)
		return
	}

	var buf [512]byte
	n, err := conn.Read(buf[0:])
	if err != nil {
		return
	}
	fmt.Println(string(buf[0:n]))
}
