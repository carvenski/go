package main

import (
    "io"
    "os"
    "log"
    "bytes"
    "io/ioutil"
    "net/http"
    "mime/multipart"
    "fmt"
)

// multipart.NewWriter
func main() {
    bodyBuffer := &bytes.Buffer{}
    bodyWriter := multipart.NewWriter(bodyBuffer)

    // 文件名
    var file_name string
    fmt.Println("Please enter the file name: ")
    fmt.Scan(&file_name)
    fmt.Println("file is uploading: ", file_name)

    fileWriter, _ := bodyWriter.CreateFormFile("file", file_name)

    file, _ := os.Open("./" + file_name)
    defer file.Close()

    // 读取传输文件
    io.Copy(fileWriter, file)

    contentType := bodyWriter.FormDataContentType()
    bodyWriter.Close()

    resp, _ := http.Post("http://21.50.131.35:37799/upload", contentType, bodyBuffer)
    // resp, _ := http.Post("http://localhost:8080/upload", contentType, bodyBuffer)
    defer resp.Body.Close()

    resp_body, _ := ioutil.ReadAll(resp.Body)

    log.Println(resp.Status)
    log.Println(string(resp_body))
    var z string
    fmt.Scan(&z)
}