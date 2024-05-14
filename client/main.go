package main

import (
	"log"
	"time"

	"github.com/gorilla/websocket"
)

var url = "ws://localhost:8080/ws"

func main() {
	conn, _, err := websocket.DefaultDialer.Dial(url, nil)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	if err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, World!")); err != nil {
		log.Fatal(err)
	}

	for {
		time.Sleep(time.Second)

		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("メッセージの種類: %d, メッセージ: %s\n", mt, message)
	}
}
