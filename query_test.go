package zormexamples

import (
	"context"
	"encoding/json"
	"fmt"
	"testing"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

//TestQueryRow 05.测试查询一个struct对象
func TestQueryRow(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	var (
		demo     models.DemoStruct     // 表
		sceneVar int               = 3 // 场景切换
	)

	switch sceneVar {
	case 1: // IN 的用法
		finder := zorm.NewSelectFinder(demo.GetTableName())
		finder.Append("WHERE id=? and active in(?)", "20210630163227149563000042432429", []int{0, 1})

		has, err := zorm.QueryRow(ctx, finder, &demo)
		if err != nil {
			t.Errorf("错误:%v", err)
		}
		t.Log(demo, has)

	case 2: // 注意目标字段count(*)的位置
		finder := zorm.NewSelectFinder(demo.GetTableName(), "count(*)")

		var count int64
		has, err := zorm.QueryRow(ctx, finder, &count)
		if err != nil {
			t.Errorf("错误:%v", err)
		}
		t.Log(count, has)

	case 3: // 注意目标字段count(*)的位置
		active := 0

		sqlStr := `
			select count(*) from t_demo
		`

		finder := zorm.NewFinder()
		finder.Append(sqlStr)
		finder.Append("where 1=?", 1)
		if active > 0 {
			finder.Append("and active =?", active)
		}

		var count int64
		has, err := zorm.QueryRow(ctx, finder, &count)
		if err != nil {
			t.Errorf("错误:%v", err)
		}
		t.Log(count, has)

	default:
		t.Errorf("场景不存在")
	}
}

//TestQueryRowMap 06.测试查询map接收结果,用于不太适合struct的场景,比较灵活
func TestQueryRowMap(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	var demo models.DemoStruct

	finder := zorm.NewSelectFinder(demo.GetTableName())
	finder.Append("WHERE id=? and active in(?)", "20210630163227149563000042432429", []int{0, 1})
	resultMap, err := zorm.QueryRowMap(ctx, finder)
	if err != nil {
		t.Errorf("错误:%v", err)
	}
	//打印结果
	fmt.Println(resultMap)
}

//TestQuery 07.测试查询对象列表
func TestQuery(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	var (
		demo     models.DemoStruct     // 表
		sceneVar int               = 3 // 场景切换
	)

	switch sceneVar {
	case 1: // 分页查询(推荐)
		finder := zorm.NewSelectFinder(demo.GetTableName())

		page := zorm.NewPage()
		page.PageNo = 1
		page.PageSize = 10

		list := make([]models.DemoStruct, 0)
		err := zorm.Query(ctx, finder, &list, page)
		if err != nil {
			t.Errorf("错误:%v", err)
		}

		data, _ := json.Marshal(list)
		fmt.Println("总条数:", page.TotalCount, "  列表:", string(data))

	case 2: // 查询所有(不推荐)
		finder := zorm.NewSelectFinder(demo.GetTableName())

		list := make([]models.DemoStruct, 0)
		err := zorm.Query(ctx, finder, &list, nil)
		if err != nil {
			t.Errorf("错误:%v", err)
		}

		data, _ := json.Marshal(list)
		fmt.Println("总条数:", len(list), "  列表:", string(data))

	case 3: // 关联表
		sqlStr := `
		select 
			t1.*,
			t2.userName,
			t2.active
		from 
			t_actor as t1 
		inner join t_demo as t2 on t2.userName = t1.stageName
		`
		active := -1

		finder := zorm.NewFinder()
		finder.Append(sqlStr)
		finder.Append("where 1=?", 1)
		if active != -1 {
			finder.Append("and t2.active =?", active)
		}
		finder.Append("order by t1.id asc")

		list := make([]models.ActorInfo, 0)
		err := zorm.Query(ctx, finder, &list, nil)
		if err != nil {
			t.Errorf("错误:%v", err)
		}

		data, _ := json.Marshal(list)
		fmt.Println("总条数:", len(list), "  列表:", string(data))

	default:
		t.Errorf("场景不存在")
	}
}

//TestQueryMap 08.测试查询map列表,用于不方便使用struct的场景,一条记录是一个map对象
func TestQueryMap(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	//构造查询用的finder
	finder := zorm.NewSelectFinder(models.DemoStructTableName) // select * from t_demo

	//创建分页对象,查询完成后,page对象可以直接给前端分页组件使用
	page := zorm.NewPage()
	page.PageNo = 1    //查询第1页,默认是1
	page.PageSize = 20 //每页20条,默认是20
	//执行查询
	listMap, err := zorm.QueryMap(ctx, finder, page)
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}
	//打印结果
	fmt.Println("总条数:", page.TotalCount, "  列表:", listMap)
}
