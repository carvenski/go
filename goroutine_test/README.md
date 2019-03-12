## difference between python coroutine & go goroutine:

#### 1.in tornado: if u don't 'yield' current pyroutine, it will run forever.
#### 2.in gevent:  if u don't use python socket, gevent will never yield current pyroutine, it will run forever too.
#### 3.in goroutine: whatever cpu compute or io block operation, goroutine will yield itself to let other goroutine run ??   

## understand channel in go:
**1.unbuffered channel = buffered channel(length=0), use for "value dependency in logic, wait for another goroutine's value".  
**2.buffered channel = Queue, use it just like queue between producer & customer.               


