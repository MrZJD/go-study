package main

import (
	"fmt"
	"log"
	"net/http"

	"golang.org/x/net/websocket"
)

// -> http headers:[Sec-WebSocket-Key]
// -> <%= Sec-WebSocket-Key %>258EAFA5-E914-47DA-95CA-C5AB0DC85B11 -> sha1 | base64 -> [Sec-WebSocket-Accept]

// go get golang.org/x/net/websocket

func main() {
	http.Handle("/", websocket.Handler(func(ws *websocket.Conn) {
		var err error

		for {
			var reply string

			if err = websocket.Message.Receive(ws, &reply); err != nil {
				fmt.Println("Can't Recevie")
				break
			}

			msg := "Recevied: " + reply
			fmt.Println("Sending to client: " + msg)

			if err = websocket.Message.Send(ws, msg); err != nil {
				fmt.Println("Can't Send")
				break
			}
		}
	}))

	if err := http.ListenAndServe(":1234", nil); err != nil {
		log.Fatal("ListenAndServe: [Error]: ", err)
	}
}
