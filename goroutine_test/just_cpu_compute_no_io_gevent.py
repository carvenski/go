from gevent import monkey; monkey.patch_all()
import gevent

def f(i):
    print('----1 ', i)
    print('----2 ', i)
    print('----3 ', i)
    print('----4 ', i)

greenlets = [gevent.spawn(f, i) for i in range(10)]
gevent.joinall(greenlets)

