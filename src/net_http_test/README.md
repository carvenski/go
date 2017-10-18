
#### msyql query or http request, actually they all use socket essentially => go's socket must like Gevent's non-block socket !!
#### that's why when http request io block, current goroutine konws how to yield CPU to let other goroutinr run.

### so how go's socket do async non-block like Gevent's socket ??
### need to do some research and see how go's socket realization !!
