package main

import "C"
import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

//export Hello
// 需要申明使用C的字符串类型,以及C和Go的类型转换
// TODO: CGo传参的话需要及时回收内存,否则内存泄漏 ?? unsafe.free(unsafe.Pointer(a))
func Hello(a *C.char) *C.char { 
	fmt.Printf("java args = %s", C.GoString(a))
	return C.CString("def")
}

//export HttpServer
// 直接启动一个httpserver,和java之间走http localhost互相调用,耗时也就增加10ms而已,可忽略.
func HttpServer() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello World From echo !")
	})
	e.Logger.Fatal(e.Start(":8002"))
}

func main() {

}

// go编译so库: go build -buildmode=c-shared -o libgo.so entry.go
