package zormexamples

import (
	"context"
	"errors"
	"testing"
	"time"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

//TestUpdateNotZeroValue 09.更新struct对象,只更新不为零值的字段.主键必须有值
func TestUpdateNotZeroValue(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//声明一个对象的指针,用于更新数据
		demo := &models.DemoStruct{}
		demo.Id = "20210630163227149563000042432429"
		demo.UserName = "冷檬"

		//更新 "sql":"UPDATE t_demo SET userName=? WHERE id=?","args":["冷檬","20210630163227149563000042432429"]
		_, err := zorm.UpdateNotZeroValue(ctx, demo)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}

//TestUpdate 10.更新struct对象,更新所有字段.主键必须有值
func TestUpdate(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//声明一个对象的指针,用于更新数据
		demo := &models.DemoStruct{}
		demo.Id = "20210630163227149563000042432429"
		demo.UserName = "陈翔"
		demo.CreateTime = time.Now()

		_, err := zorm.Update(ctx, demo)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}
}

//TestUpdateFinder 11.通过finder更新,zorm最灵活的方式,可以编写任何更新语句,甚至手动编写insert语句
func TestUpdateFinder(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	sceneVar := 3 // 场景切换

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		var finder *zorm.Finder

		switch sceneVar {
		case 1: // 更新指定字段
			finder = zorm.NewUpdateFinder(models.DemoStructTableName)
			finder.Append("userName=?,active=?", "冷檬", 1).Append("WHERE id=?", "20210630163227149563000042432429")

		case 2: // 删除操作
			finder = zorm.NewDeleteFinder(models.DemoStructTableName)
			finder.Append("where userName=?", "王炸")

		case 3: // 链式SQL更新
			finder = zorm.NewFinder().Append("UPDATE").Append(models.DemoStructTableName).Append("SET").Append("userName=?", "米线儿")
			finder.Append("where userName=?", "闰土")

		default:
			return nil, errors.New("场景不存在")
		}

		_, err := zorm.UpdateFinder(ctx, finder)
		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}

//TestUpdateEntityMap 12.更新一个EntityMap,主键必须有值
func TestUpdateEntityMap(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		//创建一个EntityMap,需要传入表名
		entityMap := zorm.NewEntityMap(models.DemoStructTableName)
		//设置主键名称
		entityMap.PkColumnName = "id"
		//Set 设置数据库的字段值,主键必须有值
		entityMap.Set("id", "20210630163227149563000042432429")
		entityMap.Set("userName", "陈翔")
		//更新 "sql":"UPDATE t_demo SET userName=? WHERE id=?","args":["TestUpdateEntityMap","20210630163227149563000042432429"]
		_, err := zorm.UpdateEntityMap(ctx, entityMap)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}
