# K8 #

## k8的各组件的高可用怎么做的
K8本身是使用etcd实现的主从选举,
也可以考虑使用etcd实现分布式锁的方案

## 一个pod内部的多个container之间的进程关系/先后顺序
每个container里面一般就只有自己的一个容器进程,
多个container之间是共享网络空间和mount空间的,可以使用localhost通信,且共享volumn.
一个pod内的多个container是按照yaml中的spec.containers指定的先后顺序启动的,
但是这并不能保证它们能按照依赖关系来正确启动顺序,还需要做一些额外的工作:
比如可以直接在各自容器的entrypoint脚本中指定等待db就绪,
也可以在pod的生命周期postStart处指定等待脚本,
K8的v18版本里还提供了Sidecar标签,它也可以指定容器需要在其他容器之前启动.
还有一种方式是使用pod的InitContainer字段来指定一些初始化等待的脚本等.

## 一个pod的生命周期/如何hook一个pod/postStart和preStop钩子
created/pending/running/unknown/failure/
在pod的各生命周期可以定义hook钩子函数,
比如创建完成后的pod可以自动上报状态等

## 内外网service的暴露方式/nginx ingress
内网的service就使用clusterIP方式内部通信即可,靠iptables作内部负载均衡,
外网访问入口需要使用service的NodePort方式,并且要配合一个外部网关来实现

## pod健康检查机制/存活探针和就绪探针
有3种检查方式: 端口检查/api检查/exec指定命令检查
可以使用exec方式来指定一段命令作为就绪探针检查,
这样就可以自定义pod的就绪状态检查逻辑了

## operator二次开发/client-go库/informer同步机制
使用operator-sdk工具来开发,基于sample-controller示例代码,
informer的内部原理就是调用K8的http api来列出资源的list,以及watch资源更新事件.

## pause容器的作用/类似init进程作用/作为1号进程/管理着其他容器进程
实现对pod级别的探针监控和进程管理.
需要一个父容器充当linux的init进程来管理多个容器进程,
K8用pause容器来作为一个pod中所有容器的父容器,这个pause容器有两个核心的功能:
第一，它提供整个pod的Linux命名空间的基础,其他容器都是共享它的命名空间,
第二，启用PID命名空间，它在每个pod中都作为PID为1的进程，并回收僵尸进程。
pause进程负责处理SIGCHID信号并处理回调,在一个子进程终止时将SIGCHLD信号发送给其父进程，
docker本身是没有这种父进程的设计的,容器里就是业务进程本身.

## pod如何重启
docker是有restart命令的,
但K8中的pod没有restart命令,只能delete+recreate删除重建

## 如何向pod注入执行命令/exec
exec接口可以用来向pod执行命令

## 关于etcd中资源的version版本号的作用
解决并发写入的冲突问题,比较版本号即可,版本小的那个就是先写入的

## pod资源限额
limits：用于限制运行时容器的最大占用资源，当容器占用资源超过limits时会被终止
requests：用于设置容器所需要申请的最小资源，如果环境资源不够，容器将无法启动

## pod自动扩缩容/弹性伸缩
水平扩展和垂直扩展两种方式
调整RC副本集里的数量,
调整pod分配的资源配额,


# linux kernel #

## 进程调度/cpu资源

## 虚拟内存/页表page table/磁盘页缓存
进程空间里看到的连续的虚拟内存地址,依靠查询页表,来映射到不连续的物理内存上,
虚拟内存本身采用分段和分页式设计,代码段和数据段和堆段和栈段,一页是4K大小,
读写虚拟地址时如果发现该地址内存还未分配,就先引发个缺页异常,然后再真正的分配物理内存,
cpu有两种运行模式,分别对应用户态和内核态,
每个进程的用户态的进程地址空间是自己独有的,而内核态的地址空间是共享的,
直接读写磁盘太慢了,所以可以使用空闲的内存来做缓存,也就是页缓存.

## 文件系统/磁盘读写

## 网络IO/异步IO/epoll机制












