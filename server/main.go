package main

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

func main() {
	handler := newHandler(websocket.Upgrader{})
	http.HandleFunc("/ws", handler)

	log.Fatal(http.ListenAndServe(":8080", nil))
}

func newHandler(upgrader websocket.Upgrader) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Fatal(err)
		}
		defer conn.Close()

		for {
			mt, message, err := conn.ReadMessage()
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("メッセージの種類: %d, メッセージ: %s\n", mt, message)

			if err = conn.WriteMessage(mt, message); err != nil {
				log.Fatal(err)
			}
		}
	}
}
