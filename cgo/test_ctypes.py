import ctypes
import time

# 1.使用ctypes调用go编译的.so库
so = ctypes.CDLL("./mylib.so")
t = time.time()
so.foo.restype = ctypes.c_longlong  # ctypes需要做一些C数据类型转换,否则结果可能不精确
r = so.foo(1, 9)
print(r)
# 耗时 2ms
print("cgo, time=%s" % (time.time()-t))  
    
t = time.time()
# 2.测试python自己的cpu计算速度
sum = 0
for i in range(1000000):
    sum += i
print(sum)
# 耗时 200ms
print("python, time=%s" % (time.time()-t))

# cpu计算型的任务速度提升了100倍

