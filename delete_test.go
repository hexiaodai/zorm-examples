package zormexamples

import (
	"context"
	"testing"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

//TestDelete 13.删除一个struct对象,主键必须有值
func TestDelete(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	sceneVar := 1 // 场景切换

	_, err := zorm.Transaction(ctx, func(ctx context.Context) (interface{}, error) {

		var err error

		switch sceneVar {
		case 1: // 删除单条数据
			demo := &models.DemoStruct{}
			demo.Id = "20210630163227149563000042432429"
			//删除 "sql":"DELETE FROM t_demo WHERE id=?","args":["20210630163227149563000042432429"]
			_, err = zorm.Delete(ctx, demo)

		case 2: // 删除多条数据
			finder := zorm.NewDeleteFinder(models.DemoStructTableName)
			finder.Append("where id in (?)", []string{"20220726134812876236000873883328"})
			_, err = zorm.UpdateFinder(ctx, finder)
		}

		//如果返回的err不是nil,事务就会回滚
		return nil, err
	})
	if err != nil { //标记测试失败
		t.Errorf("错误:%v", err)
	}

}
