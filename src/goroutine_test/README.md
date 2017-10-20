### difference between pyroutine & goroutine:
#### in tornado: if u don't 'yield' current pyroutine, it will run forever.
#### in gevent:  if u don't use python socket, gevent will never yield current pyroutine, it will run forever too.
#### in goroutine: whatever cpu compute or io block operation, goroutine will yield itself to let other goroutine run. ?? 


####understand channel in go:
1.unbuffered channel == buffered channel(length=0)

2.buffered channel == Queue


