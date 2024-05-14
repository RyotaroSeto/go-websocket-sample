package main

import (
    "fmt"
    "log"
    "net/http"

    "github.com/gorilla/websocket"
)

func main() {
    // ハンドラー関数
    upgrader := websocket.Upgrader{}
    handler := func(w http.ResponseWriter, r *http.Request) {
        // HTTP接続をWebSocket接続にアップグレード
        conn, err := upgrader.Upgrade(w, r, nil)
        if err != nil {
            log.Print(err)
            return
        }
        defer conn.Close()

        // クライアントからのメッセージをループで受信
        for {
            mt, message, err := conn.ReadMessage()
            if err != nil {
                log.Print(err)
                break
            }
            log.Printf("メッセージの種類: %d, メッセージ: %s\n", mt, message)

            // クライアントにメッセージを送信
            err = conn.WriteMessage(mt, message)
            if err != nil {
                log.Print(err)
                break
            }
        }
    }

    // HTTPサーバーを起動
    http.HandleFunc("/ws", handler)
    log.Fatal(http.ListenAndServe(":8080", nil))
}
