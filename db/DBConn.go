/*
DBConn使用:
这里主要是写点sql,稍微封装一下sql结果即可,不需要使用orm(orm还需要创建type去映射麻烦).
可参考 https://gitea.com/xorm/xorm

// 关于连接
xorm框架内部已经封装了连接的获取与释放,用户不用手动管理连接.直接使用即可.

// 查询 Query returns []map[string][]byte, QueryString returns []map[string]string, QueryInterface returns []map[string]interface{}.
results, err := engine.Query("select * from user")
results, err := engine.QueryString("select * from user where name = ?", name)
results, err := engine.QueryInterface("select * from user")

// 修改 Exec returns affected and error
affected, err := engine.Exec("update user set age = ? where name = ?", age, name)

// 事务 Use Session When You Need Transaction
session := engine.NewSession()
session.Begin()
session.Commit()
session.RollBack()
session.Close()
*/
package db

import (
	"myapp/utils"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

// 使用xorm框架,特点是简洁强大,结合使用orm和sql,顺便熟悉其源码
var DBConn *xorm.Engine

// 初始化连接池
func InitDBConn() {
	var err error
	// mysql数据库连接信息
	mysql_addr := "finedb:finedb@tcp(21.50.131.33:8080)/finedb?charset=utf8"
	DBConn, err = xorm.NewEngine("mysql", mysql_addr)
	// 设置连接池参数
	DBConn.SetMaxIdleConns(3)            //最大3个空闲连接数,保持连接的连接数
	DBConn.SetMaxOpenConns(10)           //最大10个连接
	DBConn.SetConnMaxLifetime(time.Hour) //连接的最大存活时间,<8h(mysql wait_timeout)
	if err != nil {
		utils.Logger.Error("mysql [%s] conn error: %s", mysql_addr, err)
	} else {
		err := DBConn.Ping() // 测试连接mysql
		if err != nil {
			utils.Logger.Error("mysql [%s] conn error: %s", mysql_addr, err)
		} else {
			utils.Logger.Info("mysql [%s] conn success.", mysql_addr)
			res, err := DBConn.QueryString("select version();")
			utils.Logger.Info("mysql version: %v, err: %v", res, err)
		}
	}
}
