package main

import (
    "log"
    "strconv"
)

/*
ParseInt()和ParseUint()有3个参数：

func ParseInt(s string, base int, bitSize int) (i int64, err error)
func ParseUint(s string, base int, bitSize int) (uint64, error)

bitSize参数表示转换为什么位的int/uint，有效值为0、8、16、32、64。
当bitSize=0的时候，表示转换为int或uint类型。例如bitSize=8表示转换后的值的类型为int8或uint8。

base参数表示以什么进制的方式去解析给定的字符串，有效值为0、2-36。
当base=0的时候，表示根据string的前缀来判断以什么进制去解析：
0x开头的以16进制的方式去解析，0开头的以8进制方式去解析，其它的以10进制方式解析。
*/
func main(){
    // 使用strconv.ParseInt函数可以转换二进制字符串为十进制int64
    i, err := strconv.ParseInt("11111111", 2, 64)
    log.Println(i, err)
}
