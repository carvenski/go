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
	"github.com/astaxie/beego/logs"
)

// a global package-level logger
var Logger *logs.BeeLogger

func init() {
	// log setting
	Logger = logs.NewLogger()
	// Log.Async()
	Logger.EnableFuncCallDepth(true)
	// mkdir logs && chmod 755 -R logs
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



