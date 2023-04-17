package zormexamples

import (
	"context"
	"testing"
	"time"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

//TestInsert 02.测试保存Struct对象
func TestInsert(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		var demo models.DemoStruct
		demo.UserName = "猪小明"
		demo.Password = "123456"
		demo.CreateTime = time.Now()
		demo.Active = 1
		_, err := zorm.Insert(ctx, &demo)
		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//标记测试失败
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}

//TestInsertSlice 03.测试批量保存Struct对象的Slice
//如果是自增主键,无法对Struct对象里的主键属性赋值
func TestInsertSlice(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		// 角色1
		demo1 := models.DemoStruct{
			UserName:   "吴妈",
			Active:     1,
			Password:   "123456",
			CreateTime: time.Now(),
		}

		// 角色2
		demo2 := models.DemoStruct{
			UserName:   "腿腿",
			Active:     1,
			Password:   "123456",
			CreateTime: time.Now(),
		}

		demoSlice := make([]zorm.IEntityStruct, 0)
		demoSlice = append(demoSlice, &demo1, &demo2)

		//批量保存对象,如果主键是自增,无法保存自增的ID到对象里.
		_, err := zorm.InsertSlice(ctx, demoSlice)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//标记测试失败
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}

//TestInsertEntityMap 04.测试保存EntityMap对象,用于不方便使用struct的场景,使用Map作为载体
func TestInsertEntityMap(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {
		var demo models.DemoStruct
		entityMap := zorm.NewEntityMap(demo.GetTableName())

		//Set 设置数据库的字段值
		//如果主键是自增或者序列,不要entityMap.Set主键的值
		entityMap.Set("id", zorm.FuncGenerateStringID(ctx))
		entityMap.Set("userName", "王炸")
		entityMap.Set("password", "123456")
		entityMap.Set("createTime", time.Now())
		entityMap.Set("active", nil)

		//执行
		_, err := zorm.InsertEntityMap(ctx, entityMap)

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	//标记测试失败
	if err != nil {
		t.Errorf("错误:%v", err)
	}
}
