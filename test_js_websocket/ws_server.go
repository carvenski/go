package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
)

func main() {

    r := gin.Default()

    r.GET("/ws", func(c *gin.Context) {
        wshandler(c)
    })

    r.Run("0.0.0.0:8888")
}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// websocket协议也是可以基于http协议实现的，所以可以直接接入到http handler里面。
func wshandler(c *gin.Context) {
    w := c.Writer
    r := c.Request

    fmt.Println("=== ws conn ok ===")
    fmt.Println(r.Header)

    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    conn.WriteMessage(websocket.BinaryMessage, []byte(" from ws server."))
    conn.Close()
   
}

