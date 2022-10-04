import web
import time

urls = (
    '/(.*)', 'hello'
)
app = web.application(urls, globals())

class hello:
    def GET(self, path):
        if not path:
            path = 'World'
        time.sleep(3)
        return 'Hello, ' + path + '!'

if __name__ == "__main__":
    app.run()
