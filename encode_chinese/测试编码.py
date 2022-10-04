a = "测试"

def str_bytes_encode(str, encode):
    l = list( a.encode(encode) )
    print("- 中文 转 bytes: "+encode)
    print(l)
    print("- bytes 转 中文: "+encode)
    print( bytes(l).decode(encode) )

str_bytes_encode("测试", "utf8")
str_bytes_encode("测试", "gbk")

print("\n字符串是相同的,而按照不同的编码编出的bytes是不同的.\n编码后的字节数组中的数字是不同的.")
