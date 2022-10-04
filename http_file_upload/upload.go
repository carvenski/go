package main 
import(
    "os"
    "fmt"
    "net/http"
    "log"
    "text/template"
    "crypto/md5"
    "time"
    "io"
    "strconv"
)

func upload(w http.ResponseWriter, r *http.Request){
    fmt.Println("method", r.Method) //获得请求的方法
    
    if r.Method == "GET"{ //
        html := `
<html>
<head>
<title>上传文件</title>
</head>
<body>
<form enctype="multipart/form-data" action="http://192.168.44.130:9090/upload" method="post">
    <input type="file" name="uploadfile" />
    <input type="hidden" name="token" value="{{.}}" />
    <input type="submit" value="upload" />
</form>
</body>
</html>
`
        crutime := time.Now().Unix()
        h := md5.New()
        io.WriteString(h, strconv.FormatInt(crutime, 10))
        token := fmt.Sprintf("%x", h.Sum(nil))

        t := template.Must(template.New("test").Parse(html))
        t.Execute(w, token)
    }else{
        //表示maxMemory,调用ParseMultipart后，上传的文件存储在maxMemory大小的内存中，如果大小超过maxMemory，剩下部分存储在系统的临时文件中
        r.ParseMultipartForm(32 << 20)
        //根据input中的name="uploadfile"来获得上传的文件句柄
        file, handler, err := r.FormFile("uploadfile")
        if err != nil{
            fmt.Println(err)
            return
        }
        defer file.Close()
        //得到上传文件的Header和文件名
        fmt.Fprintf(w, "%v,%s", handler.Header, handler.Filename)
        
        //然后打开该文件
        openFile, err := handler.Open()
        if err != nil {
            fmt.Println(err)
            return
        }
        dst, _ := os.Create("./" + handler.Filename + ".upload")
        defer dst.Close()
        io.Copy(dst, openFile)
    }
}

func main() {
    http.HandleFunc("/upload", upload)         //设置访问的路由
    err := http.ListenAndServe(":9090", nil) //设置监听的端口
    if err != nil{
        log.Fatal("ListenAndServe : ", err)
    }
}