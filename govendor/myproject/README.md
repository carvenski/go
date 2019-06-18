#### govendor使用

```
安装 go get -u -v github.com/kardianos/govendor

# 进入到项目目录
cd /home/gopath/src/mytool

# 初始化vendor目录
govendor init

# govendor会将GOPATH中本工程使用到的依赖包自动移动到vendor目录中
# (说明：如果本地GOPATH没有依赖包，需要先go get安装好相应的依赖包,govendor只是生成依赖文件并拷贝包到vendor/而已)
govendor add +external (或使用缩写： govendor add +e)

# 其他人有了vendor.json文件后,则可以本地拉取所有包到vendor
(这样就不需要把vendor里面的包文件也上传了,当然你直接上传vendor也行)
govendor sync


```
