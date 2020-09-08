### windows下编译成linux的可执行文件
```
设置3个环境变量即可(注意几个环境变量后面不能有空格):
SET CGO_ENABLED=0
SET GOOS=linux
SET GOARCH=amd64

然后直接编译:
go build main.go
```




