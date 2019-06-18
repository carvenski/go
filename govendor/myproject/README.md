#### govendor使用

```
go get -u -v github.com/kardianos/govendor

#进入到项目目录
cd /home/gopath/src/mytool

#初始化vendor目录
govendor init

#将GOPATH中本工程使用到的依赖包自动移动到vendor目录中
#说明：如果本地GOPATH没有依赖包，需要先go get安装好相应的依赖包,govendor只是生成依赖文件并拷贝下包而已
govendor add +external (或使用缩写： govendor add +e)

```
