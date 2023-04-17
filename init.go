package zormexamples

import (
	"log"

	"gitee.com/chunanyong/zorm"
	//util "gitee.com/chunanyong/zorm-examples/util"
	// _ "github.com/go-sql-driver/mysql"
	_ "kingbase.com/gokb"
)

var DbDao *zorm.DBDao

// 01.初始化DBDao
func init() {
	dbConfig := zorm.DataSourceConfig{
		//连接数据库DSN
		// DSN: "dm://SYSDBA:SYSDBA@127.0.0.1:5236", //DM
		// dsn := "host=127.0.0.1 user=SYSTEM password=123456 dbname=SAMPLES port=54321 sslmode=disable"
		DSN: "host=10.31.203.93 user=keycloak password=7mB6sZzwiB dbname=keycloak port=54321 sslmode=disable",
		// DSN: "keycloak:7mB6sZzwiB@tcp(10.31.203.93:54321)/keycloak?charset=utf8", //MySQL
		//数据库类型
		// DriverName: "dm",
		// DBType:     "dm",
		//sql.Open(DriverName,DSN) DriverName就是驱动的sql.Open第一个字符串参数,根据驱动实际情况获取
		DriverName: "kingbase",
		Dialect:    "kingbase",
		//SlowSQLMillis 慢sql的时间阈值,单位毫秒.小于0是禁用SQL语句输出;等于0是只输出SQL语句,不计算执行时间;大于0是计算SQL执行时间,并且>=SlowSQLMillis值
		SlowSQLMillis: 0,
		//最大连接数 默认50
		MaxOpenConns: 0,
		//最大空闲数 默认50
		MaxIdleConns: 0,
		//连接存活秒时间. 默认600
		ConnMaxLifetimeSecond: 0,
		//事务隔离级别的默认配置,默认为nil
		DefaultTxOptions: nil,
	}

	var err error
	DbDao, err = zorm.NewDBDao(&dbConfig)
	if err != nil {
		log.Fatalf("数据库连接异常 %v", err)
	}
	log.Println("数据库连接成功")

	// 达梦TEXT转换成string类型插件,dialectColumnType 值是 Dialect.字段类型 ,例如 dm.TEXT
	// 一般是放到init方法里进行注册
	// zorm.RegisterCustomDriverValueConver("dm.TEXT", util.CustomDMText{})
}
