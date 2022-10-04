## java调用go 技术方案 
*不再写java,全部转化成写go => 我是py/go语言专家*

### 场景一: java调用go
- java使用JNA可直接调用go编译的.so库，这就是java调用原生的c库函数的推荐方式。
性能高，需注意及时释放JNA申请的Pointer/Memory内存，防止内存泄漏bug。

## 场景二: go调用java
某个功能的库只有jar包版本，必须调用java。
例如太保这需要调用一个p13密码生成的函数库，只提供一个jar包。
- 只能是走微服务调用,启动一个本地httpserver,http调用。必须要启动两个独自的进程。

## 场景三: java调go 同时 go调java
- 多线程, jna, 微服务, http调用
在java里启动2个线程(多核运行)，一个运行java，一个使用jna运行go编译的.so库。
分别启动httpserver各占用一个端口8001,8002,互相之间使用http互调。
(这里补充一下:不需要http调用,直接开启一个http代理到go服务即可)
从外部看只需要启动一个jvm进程即可，里面靠2个线程来分开即可。

## http代理
可以使用jetty的ProxyServlet作为http代理,      
更推荐使用go的httputil.NewSingleHostReverseProxy(url)来做代理.      


