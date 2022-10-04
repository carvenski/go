package utils

import "net/http"

/*
net/http.Client里的TransPort里是有个连接池的:
在处理完一次http请求后，会立即把当前连接放到连接池中,
每个client的连接池结构是这样的: idleConn map[connectMethodKey][]*persistConn
其中connectMethodKey的值就是client连接的server的host值,
map的值是一个*persistConn类型的slice结构,就是存放连接的地方,
slice的长度由MaxIdleConnsPerHost这个值指定的,
当我们不设置这个值的时候就取默认的设置: const DefaultMaxIdleConnsPerHost = 2
从上面的介绍可知: net/http默认就是连接复用的,对于每个server默认的连接池大小是2
*/
// 全局的http.Client对象,所有的http请求使用它来发送
var httpClient *http.Client = &http.Client{}
