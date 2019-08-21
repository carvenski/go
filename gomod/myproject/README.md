#### go mod使用
```
require: go version > go1.11

示例:
cd myproject/
go mod init myproject  //初始化生成go.mod文件
go mod tidy            //拉取缺少的模块，移除不用的模块

go run/build main.go   //此时项目依赖已经全部下载好了
```

