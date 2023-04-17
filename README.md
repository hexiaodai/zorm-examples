## zorm-examples
zorm的示例代码

## 要求
+ Go 1.13 or higher.
+ Mysql (5.7+)

## 安装
导入数据，见文件夹script里的t_actor.sql、t_demo.sql文件

## 测试
输入命令
```
cd zorm-examples

go test -v -run TestQuery
```
测试结果
```
wuxian@wxdeMacBook-Air zorm-examples % go test -v -run TestQuery
2022/07/26 08:17:20 /Users/wuxian/Documents/opensource/zorm-examples/init.go:41: 数据库连接成功
=== RUN   TestQueryRow
2022/07/26 08:17:20 /Users/wuxian/Documents/opensource/zorm-examples/query_test.go:54: sql:  
                        select count(*) from t_demo
                 where 1=? ,args: [1]
--- PASS: TestQueryRow (0.02s)
    query_test.go:58: 6 true
=== RUN   TestQuery
2022/07/26 08:17:20 /Users/wuxian/Documents/opensource/zorm-examples/query_test.go:121: sql:  
                select 
                        t1.*,
                        t2.userName,
                        t2.active
                from 
                        t_actor as t1 
                inner join t_demo as t2 on t2.userName = t1.stageName
                 where 1=? order by t1.id asc ,args: [1]
总条数: 6   列表: [{"Id":"","StageName":"陈翔","RealName":"陈翔","Company":"陈翔六点半","CreateTime":"2022-07-26T07:41:48Z","DemoStructId":"1","DemoStructUserName":"陈翔","DemoStructPassword":"","DemoStructCreateTime":"0001-01-01T00:00:00Z","DemoStructActive":0},{"Id":"","StageName":"妹总","RealName":"应宝林","Company":"陈翔六点半","CreateTime":"2022-07-26T07:44:20Z","DemoStructId":"2","DemoStructUserName":"妹总","DemoStructPassword":"","DemoStructCreateTime":"0001-01-01T00:00:00Z","DemoStructActive":0},{"Id":"","StageName":"球球","RealName":"纪文君","Company":"陈翔六点半","CreateTime":"2022-07-26T07:45:22Z","DemoStructId":"3","DemoStructUserName":"球球","DemoStructPassword":"","DemoStructCreateTime":"0001-01-01T00:00:00Z","DemoStructActive":0},{"Id":"","StageName":"毛台","RealName":"邰光远","Company":"陈翔六点半","CreateTime":"2022-07-26T07:48:51Z","DemoStructId":"4","DemoStructUserName":"毛台","DemoStructPassword":"","DemoStructCreateTime":"0001-01-01T00:00:00Z","DemoStructActive":0},{"Id":"","StageName":"闰土","RealName":"李闰刚","Company":"陈翔六点半","CreateTime":"2022-07-26T07:50:08Z","DemoStructId":"5","DemoStructUserName":"闰土","DemoStructPassword":"","DemoStructCreateTime":"0001-01-01T00:00:00Z","DemoStructActive":0},{"Id":"","StageName":"蘑菇头","RealName":"黄晓飞","Company":"陈翔六点吧","CreateTime":"2022-07-26T07:51:02Z","DemoStructId":"6","DemoStructUserName":"蘑菇头","DemoStructPassword":"","DemoStructCreateTime":"0001-01-01T00:00:00Z","DemoStructActive":0}]
--- PASS: TestQuery (0.01s)
PASS
ok      gitee.com/chunanyong/zorm-examples      0.054s
wuxian@wxdeMacBook-Air zorm-examples % 
```
