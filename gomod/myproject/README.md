## go mod使用姿势
```
require: go version > go1.11

GOPATH始终需要设置的,
只是使用了go mod后,项目不需要再放在GOPATH/src下了.


示例:
go mod init myproject  //新建myproject项目,并生成go.mod文件
go mod tidy/download   //清理/下载依赖库,放到GOPATH/pkg/mod下
go mod graph           //打印全部依赖库
go run/build main.go   //默认会自动拉取项目依赖后再运行
go list -m -versions 包名 //列出包可下载的所有版本

修改依赖库的版本:
直接修改go.mod文件中对应库的版本号即可.
```

### 关于/vN后缀
```
go mod支持包的不同版本的写法: /vN后缀
以echo包举例:
"github.com/labstack/echo"     这个路径只支持到v3.3.10
"github.com/labstack/echo/v4"  加了v4后缀的这个路径才能支持v4.0.0以上的版本 
这算个go mod的坑吧...
```
