package main

import (
    "fmt"
    "github.com/gin-gonic/gin"
    "github.com/gorilla/websocket"
    "encoding/json"
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

type Msg struct {
    A int
    B string
}

var wsupgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
}

// websocket协议也是可以基于http协议实现的，所以可以直接接入到http handler里面。
func wshandler(c *gin.Context) {
    w := c.Writer
    r := c.Request
    conn, err := wsupgrader.Upgrade(w, r, nil)
    if err != nil {
        fmt.Println("Failed to set websocket upgrade: %+v", err)
        return
    }

    for {
        _, msg, err := conn.ReadMessage()
        if err != nil {
            break
        }
        var m Msg
        err = json.Unmarshal(msg, &m)
        if err != nil {
            fmt.Println("Json Unmarshal fail: %+v", err)
        } else {
            fmt.Println("Json Unmarshal ok: %+v", m)
        }
        conn.WriteMessage(websocket.BinaryMessage, []byte(string(msg) + " from ws server."))
    }
}

func wsclienthandler(c *gin.Context) {
    wsUrl := "ws://localhost:8080" + "/ws"
    
    // websocket client端的使用类似socket client    
    conn, _, err := websocket.DefaultDialer.Dial(wsUrl, nil)
    if err != nil {
        fmt.Println("Error connecting to Websocket Server:", err)
    }
    defer conn.Close()

    msg, _ := json.Marshal(Msg{1, "b"})
    err = conn.WriteMessage(websocket.BinaryMessage, msg)
    if err != nil {
        fmt.Println("Error during writing to websocket:", err)
        return
    }
    _, msg, err = conn.ReadMessage()
    if err != nil {
        fmt.Println("Error in receive:", err)
        return
    }

    c.String(200, "ws client got response: "+string(msg))

}
