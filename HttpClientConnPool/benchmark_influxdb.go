package main
 
import (
	"os"
	"log"
	"time"
	"strconv"
	"math/rand"

	// this repository api has changed, use git branch=1.8, shit.
	"github.com/influxdata/influxdb/client/v2"
)
 
const (
	addr = "http://10.10.25.28:8086"
	db = "benchmark"
	username = ""
	password = ""
	precision = "s"
)

func uploadMetrics(node int, start chan bool, conn *client.Client)  {	
	nodeID := "fku60798840306e5b2e48272f7-node" + strconv.Itoa(node)
	for {
		// waiting signal to start upload
		<- start
		log.Printf("node %d is writing.", node)

		// BatchPoints
		bp, err := client.NewBatchPoints(client.BatchPointsConfig{
			Database:  db,
			Precision: precision,
		})
		if err != nil {
			log.Println(err)
		}
		
		t := time.Now()

		// table basic:	
		metrics := []string{"bCpuUsage","bCpuIdle","bLoadAvg1","bLoadAvg5","bLoadAvg15","bLoadAvg15","bTcpTotal","bTcpEstablished","bTcpSynSent","bTcpSynRecv","bTcpFinWait1","bTcpFinWait2","bTcpTimeWait","bTcpClose","bTcpCloseWait","bTcpLastAck","bTcpListen","bTcpClosing","bMemUsed","bMemFree","bMemBuffers","bMemCached","bMemRealFree","bMemRealUsed","bMemUsage","xbCpuUsage","xbCpuIdle","xbLoadAvg1","xbLoadAvg5","xbLoadAvg15","xbLoadAvg15","xbTcpTotal","xbTcpEstablished","xbTcpSynSent","xbTcpSynRecv","xbTcpFinWait1","xbTcpFinWait2","xbTcpTimeWait","xbTcpClose","xbTcpCloseWait","xbTcpLastAck","xbTcpListen","xbTcpClosing","xbMemUsed","xbMemFree","xbMemBuffers","xbMemCached","xbMemRealFree","xbMemRealUsed","xbMemUsage"}
		for _, metric := range metrics {
				tags := map[string]string{
					"uuid": nodeID,
					"metric": metric,
				}
			fields := map[string]interface{}{"value": rand.Intn(1000)}
			pt, err := client.NewPoint("basic", tags, fields, t)
			if err != nil {
				log.Println(err)
			}
			bp.AddPoint(pt)
		}

		// table interface:		
		metrics = []string{"iTotalBitInSec","iTotalPacketInSec","iTotalErrorInSec","iTotalDropInSec","iTotalBitOutSec","iTotalPacketOutSec","iTotalErrorOutSec","iTotalDropOutSec","xiTotalBitInSec","xiTotalPacketInSec","xiTotalErrorInSec","xiTotalDropInSec","xiTotalBitOutSec","xiTotalPacketOutSec","xiTotalErrorOutSec","xiTotalDropOutSec"}
		for _, metric := range metrics {
				tags := map[string]string{
					"uuid": nodeID,
					"metric": metric,
				}
			fields := map[string]interface{}{"value": rand.Intn(1000)}
			pt, err := client.NewPoint("interface", tags, fields, t)
			if err != nil {
				log.Println(err)
			}
			bp.AddPoint(pt)
		}

		// table disk:		
		metrics = []string{"dReadCompleteSpeed","dWriteCompleteSpeed","dReadMergeSpeed","dWriteMergeSpeed","dReadByteSpeed","dWriteByteSpeed","dUtil","dReadUtil","dWriteUtil","dAvailable","dUsag","xdReadCompleteSpeed","xdWriteCompleteSpeed","xdReadMergeSpeed","xdWriteMergeSpeed","xdReadByteSpeed","xdWriteByteSpeed","xdUtil","xdReadUtil","xdWriteUtil","xdAvailable","xdUsag"}
		for _, metric := range metrics {
				tags := map[string]string{
					"uuid": nodeID,
					"metric": metric,
				}
			fields := map[string]interface{}{"value": rand.Intn(1000)}
			pt, err := client.NewPoint("disk", tags, fields, t)
			if err != nil {
				log.Println(err)
			}
			bp.AddPoint(pt)
		}

		// table proc:		
		metrics = []string{"prOnline","prMemUsed","prCpuUsage","prThreadNum","prOpenFilesCount","xprOnline","xprMemUsed","xprCpuUsage","xprThreadNum","xprOpenFilesCount"}
		for _, metric := range metrics {
				tags := map[string]string{
					"uuid": nodeID,
					"metric": metric,
				}
			fields := map[string]interface{}{"value": rand.Intn(1000)}
			pt, err := client.NewPoint("proc", tags, fields, t)
			if err != nil {
				log.Println(err)
			}
			bp.AddPoint(pt)
		}

		// write BatchPoints with 100 points
		if err := (*conn).Write(bp); err != nil {
			log.Printf("node %d write fail.       err => %s", node, err)
		} else {
			log.Printf("node %d write success.", node)
		}
	}
}

// influxdb go client has a http connection pool inside !!
/*
// influxdb conn pool, 1 conn for 1000 nodes when concurrency
var connPool = map[int]*client.Client{}
func getConn(i int) *client.Client {
	k := i/1000
	conn, ok := connPool[k]
	// lazy load conn
	if !ok {
		c, err := client.NewHTTPClient(client.HTTPConfig{
				Addr: addr,
				Username: username,
				Password: password,
			})
		if err != nil {
			log.Panicln(err)
		}
		log.Printf("influxdb conn %d to: %s successfully.", k, addr)
		connPool[k] = &c
		return connPool[k]
	}
	return conn	
}
*/

// 如何使用一个chan实现 广播事件 ?? 而不是N个chan
func main() {
	conn, err := client.NewHTTPClient(client.HTTPConfig{
			Addr: addr,
			Username: username,
			Password: password,
		})
	if err != nil {
		log.Panicln(err)
	}
	log.Printf("influxdb conn to: %s successfully.", addr)

	defer func ()  {
		if r := recover(); r != nil {
			log.Println("ERROR => ", r)
		}
		// close all conn when exit
		// for _, conn := range connPool {
		// 	(*conn).Close()
		// }
		conn.Close()
		log.Println("benchmark exit.")
		os.Exit(0)
	}()
	if len(os.Args) != 3 {
		panic("please pass frequency time(s) & number of nodes.")
	}
	interval, _ := strconv.Atoi(os.Args[1])  // frequency time(s)
	count, _ := strconv.Atoi(os.Args[2])  // number of nodes=goroutines		   	
	
	log.SetFlags(log.Lshortfile | log.LstdFlags)	

	// write metrics data into database benchmark every 5s	
	chans := []chan bool{}
	// start all goroutines
	for i := 0; i < count; i++ {
		start := make(chan bool, 0)
		// go uploadMetrics(i, start, getConn(i))
		go uploadMetrics(i, start, &conn)
		chans = append(chans, start)
	}
	log.Println("benchmark is started.")
	log.Printf("simulate nodes count => %d.", count)
	for {
		time.Sleep(time.Second*time.Duration(interval))
		for _, ch := range chans {
			ch <- true   // send start signal every 5s
		}
	}
	
// CQ test
CQ := `
CREATE CONTINUOUS QUERY benchmark_m ON benchmark BEGIN SELECT mean(id), max(id), min(id) INTO benchmark_m FROM benchmark GROUP BY host, time(30s) END

CREATE CONTINUOUS QUERY benchmark_cq_m ON benchmark RESAMPLE FOR 2m BEGIN SELECT max(*),min(*),mean(*),sum(*),count(*) INTO benchmark_cq.autogen.:MEASUREMENT FROM benchmark.autogen./.*/ GROUP BY time(1m),* END
`
_ = CQ
}


/* ERROR:
1.
benchmark_influxdb.go:102: node 36416 write fail. 
err=Post http://10.10.25.28:8086/write?consistency=&db=benchmark&precision=s&rp=: dial tcp 10.10.25.28:8086: socket: too many open files

进程打开的文件描述符太多,超过了系统限制,需要ulimt -n 10000来调大fd限制数. 但是这是client还是server报的错误?
最终发现: 
client端的ulimt已经调到了10000,还是报错,
再把influxdb server端的max-connection-limit参数设置为10000才行,这个值是influxdb server端允许的最大连接数.
因为每个socket连接是需要client-server两端的.

2.
benchmark_influxdb.go:102: node 33378 write fail. 
err={"error":"partial write: max-series-per-database limit exceeded: (1000000) dropped=49"}

By default max-series-per-database is set to one million. Changing the setting to 0 allows an unlimited number of series per database.
*/


/*******************************************************************************************
// 参考 http://oohcode.com/2018/06/01/golang-http-client-connection-pool/
设置 go influxdb client conn pool size:
tr := &http.Transport{
        TLSClientConfig: &tls.Config{
            InsecureSkipVerify: conf.InsecureSkipVerify,
        },
        //Proxy: conf.Proxy,
        //Proxy: ProxyFromEnvironment, //代理使用
        DialContext: (&net.Dialer{
            Timeout:   30 * time.Second, //连接超时时间
            KeepAlive: 30 * time.Second, //连接保持超时时间
            DualStack: true,             //
        }).DialContext,
        MaxIdleConns:          2000, //client对与所有host最大空闲连接数总和
        MaxIdleConnsPerHost:   1000, //client对与一台host最大空闲连接数总和
        IdleConnTimeout:       90 * time.Second, //空闲连接在连接池中的超时时间
        TLSHandshakeTimeout:   10 * time.Second, //TLS安全连接握手超时时间
        ExpectContinueTimeout: 1 * time.Second,  //发送完请求到接收到响应头的超时时间
    }
*******************************************************************************************/

/*
use htop/iotop/iftop/iostat
watch disk IO: iostat -d sda -m 5  100

这是分钟粒度的测试数据
测试influxdb(10.10.25.28) 每台node每5s写入100条point (参考zabbix性能指标)
nodes     cpu     mem      net      
1000      22%     5.1G     14Mbps
2000      43%     8.85G    27Mbps    
3000      55%     12.4G    40Mbps
5000      65%     21G      58Mbps  
分钟粒度的CQ没问题,整点时触发的CQ导致大量磁盘读,内存满31G,cpu全是iowait,占用80%


10.30 开跑 1000 node

*/



