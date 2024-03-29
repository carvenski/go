/*
see:
https://beego.me/docs/module/logs.md
https://www.jianshu.com/p/b651801178d2

logs.SetLogger(logs.AdapterFile, `{"filename":"test.log"}`)

Parameters:
filename: Save to filename.
maxlines: Maximum lines for each log file, 1000000 by default.
maxsize: Maximum size of each log file, 1 << 28 or 256M by default.
daily: If log rotates by day, true by default.
maxdays: Maximum number of days log files will be kept, 7 by default.
rotate: Enable logrotate or not, true by default.
level: Log level, Trace by default.
perm: Log file permission
*/
package log

import (
    "os"
    "path/filepath"
	"github.com/astaxie/beego/logs"
)

// 全局的Logger对象,引入log包并使用,log.Logger.info("xxx");
var Logger *logs.BeeLogger

// init Logger
func init() {
	Logger = logs.NewLogger()
	// Logger.Async()
	Logger.EnableFuncCallDepth(true)
	// mkdir logs/ && chmod 755 -R logs/
    err := os.MkdirAll(filepath.Dir("logs/test.log"), 0755)
    if err != nil { 
        panic("create logs/ dir error !") 
    }
    // Logger setting, Rolling File
	Logger.SetLogger(logs.AdapterFile,
		`{"filename": "logs/test.log",
		  "level": 7,
		  "maxlines": 10000,
		  "daily": true,
		  "maxdays": 7,
		  "rotate": true,
		  "color": true
      }`)
}


