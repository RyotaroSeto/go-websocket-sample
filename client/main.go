package main

import (
    "fmt"
    "log"
    "net/http"
    "time"

    "github.com/gorilla/websocket"
)

func main() {
    // ダイヤルしてWebSocket接続を確立
    dialer := websocket.DefaultDialer
    url := "ws://localhost:8080/ws"
    conn, _, err := dialer.Dial(url, nil)
    if err != nil {
        log.Fatal(err)
    }
    defer conn.Close()

    // サーバーにメッセージを送信
    err = conn.WriteMessage(websocket.TextMessage, []byte("Hello, World!"))
    if err != nil {
        log.Fatal(err)
    }

    // サーバーからのメッセージを1秒ごとに受信
    for {
        time.Sleep(time.Second)

        mt, message, err := conn.ReadMessage()
        if err != nil {
            log.Fatal(err)
        }
        log.Printf("メッセージの種類: %d, メッセージ: %s\n", mt, message)
    }
}
