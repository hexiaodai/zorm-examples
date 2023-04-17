package zormexamples

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

//TestOther 16.其他的一些说明.非常感谢您能看到这一行
func TestOther(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	//场景1.多个数据库.通过对应数据库的dbDao,调用BindContextDBConnection函数,把这个数据库的连接绑定到返回的ctx上,然后把ctx传递到zorm的函数即可
	//也可以重写FuncReadWriteStrategy函数,通过ctx设置不同的key,返回指定数据库的DBDao
	newCtx, err := DbDao.BindContextDBConnection(ctx)
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

	finder := zorm.NewFinder().Append("SELECT * FROM " + models.DemoStructTableName) // select * from t_demo
	//把新产生的newCtx传递到zorm的函数
	list, _ := zorm.QueryMap(newCtx, finder, nil)
	fmt.Println(list)

	//场景2.单个数据库的读写分离.设置读写分离的策略函数.
	zorm.FuncReadWriteStrategy = myReadWriteStrategy

	//场景3.如果是多个数据库,每个数据库还读写分离,按照 场景1 处理.
	//也可以重写FuncReadWriteStrategy函数,通过ctx设置不同的key,返回指定数据库的DBDao

}

//myReadWriteStrategy 数据库的读写分离的策略 rwType=0 read,rwType=1 write
//也可以通过ctx设置不同的key,返回指定数据库的DBDao
func myReadWriteStrategy(ctx context.Context, rwType int) (*zorm.DBDao, error) {
	//根据自己的业务场景,返回需要的读写dao,每次需要数据库的连接的时候,会调用这个函数
	// if rwType == 0 {
	// 	return dbReadDao
	// }
	// return dbWriteDao

	return DbDao, nil
}
