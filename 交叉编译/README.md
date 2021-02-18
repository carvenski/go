### windows下编译成linux的可执行文件
```
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build main.go
```

### Linux 下编译 Mac 和 Windows 64位可执行程序
```
CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build main.go
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build main.go
```

*所以其实就是修改 GOOS 这个变量指向需要编译的平台即可。*


