
### msyql or http request, actually they all use socket essentially => go's socket must like Gevent's non-block socket !!
### that's why when http request io block, current goroutine konws how to yield CPU to let other goroutinr run.

