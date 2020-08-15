## go mod使用姿势
```
require: go version > go1.11

示例:
go mod init myproject  //新建myproject项目,并生成go.mod文件
go mod tidy/download   //清理/下载依赖库,放到GOPATH/pkg/mod下
go mod graph           //打印全部依赖库
go run/build main.go   //默认会自动拉取项目依赖后再运行
go list -m -versions 包名 //列出包可下载的所有版本

修改依赖库的版本:
直接修改go.mod文件中对应库的版本号即可.
```

