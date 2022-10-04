package main

import (
    "io"
    "os"
    "fmt"
    "io/ioutil"
    "net/http"
)

// MultipartReader
func uploadHandler(w http.ResponseWriter, r *http.Request) {
    reader, err := r.MultipartReader()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }

    for {
        part, err := reader.NextPart()
        if err == io.EOF {
            break
        }

        fmt.Printf("FileName=[%s], FormName=[%s]\n", 
            part.FileName(), part.FormName())

        if part.FileName() == "" {  // this is FormData
            data, _ := ioutil.ReadAll(part)
            fmt.Printf("FormData=[%s]\n", string(data))
        } else {    // This is FileData
            dst, _ := os.Create("./" + part.FileName())
            defer dst.Close()
            // 写入传输文件
            io.Copy(dst, part)
        }
    }
}

func main() {
    http.HandleFunc("/upload", uploadHandler)
    http.ListenAndServe(":8080", nil)
}