package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

func main() {

    r := gin.Default()

    r.LoadHTMLFiles("index.html")
    r.GET("/", func(c *gin.Context) {
        c.HTML(200, "index.html", nil)
    })

    r.GET("/ws", func(c *gin.Context) {
        wshandler(c)
    })

    r.GET("/ws/client", func(c *gin.Context) {
        wsclienthandler(c)
    })

    r.Run("0.0.0.0:8080")
}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

func wshandler(c *gin.Context) {
    w := c.Writer
    r := c.Request
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    for {
        t, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        conn.WriteMessage(t, []byte(string(msg) + " from ws server."))
    }
}

func wsclienthandler(c *gin.Context) {
    wsUrl := "ws://localhost:8080" + "/ws"
    conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
    if err != nil {
        fmt.Println("Error connecting to Websocket Server:", err)
    }
    defer conn.Close()

    err = conn.WriteMessage(websocket.TextMessage, []byte("Hello from GolangDocs!"))
    if err != nil {
        fmt.Println("Error during writing to websocket:", err)
        return
    }
    _, msg, err := conn.ReadMessage()
    if err != nil {
        fmt.Println("Error in receive:", err)
        return
    }

    c.String(200, "ws client got response: "+string(msg))

}
