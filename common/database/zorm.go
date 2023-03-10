package database

import (
	"context"
	"gitee.com/chunanyong/zorm"
	_ "github.com/go-sql-driver/mysql"
	"github.com/zeromicro/go-zero/core/logx"
)

// InitZorm 初始化数据库
func InitZorm(dataSource string) {
	logx.Infof("数据库 url: %s", dataSource)
	var err error
	dbConfig := &zorm.DataSourceConfig{
		//DSN 数据库的连接字符串,parseTime=true会自动转换为time格式,默认查询出来的是[]byte数组
		DSN: dataSource,
		//sql.Open(DriverName,DSN) DriverName就是驱动的sql.Open第一个字符串参数,根据驱动实际情况获取
		DriverName:            "mysql", //数据库驱动名称
		Dialect:               "mysql", //数据库类型
		SlowSQLMillis:         1,       //慢sql的时间阈值,单位毫秒.小于0是禁用SQL语句输出;等于0是只输出SQL语句,不计算执行时间;大于0是计算SQL执行时间,并且>=SlowSQLMillis值
		MaxOpenConns:          0,       //数据库最大连接数,默认50
		MaxIdleConns:          0,       //数据库最大空闲连接数,默认50
		ConnMaxLifetimeSecond: 0,       //连接存活秒时间. 默认600(10分钟)后连接被销毁重建.
		//避免数据库主动断开连接,造成死连接.MySQL默认wait_timeout 28800秒(8小时)
		DefaultTxOptions: nil, //事务隔离级别的默认配置,默认为nil
	}

	_, err = zorm.NewDBDao(dbConfig)

	// 把执行的sql及参数和时间写入log
	zorm.FuncPrintSQL = func(ctx context.Context, sqlstr string, args []interface{}, execSQLMillis int64) {
		logx.Infof("sql: %s, args: %+v, execSQLMillis: %d", sqlstr, args, execSQLMillis)
	}

	if err != nil {
		logx.Errorf("数据库连接异常 %v", err)
		panic(err)
	} else {
		logx.Infof("zorm 连接数据库成功 ...")
	}
}
