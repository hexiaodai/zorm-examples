package zormexamples

import (
	"context"
	"fmt"
	"testing"

	"gitee.com/chunanyong/zorm"
	"gitee.com/chunanyong/zorm-examples/models"
)

/*

-- 定义存储过程
DELIMITER $$
DROP PROCEDURE IF EXISTS testproc;
CREATE PROCEDURE testproc(IN id VARCHAR(255))
BEGIN
	SELECT * FROM t_demo WHERE id = id;
	COMMIT;
END $$

-- 调用存储过程
-- DELIMITER ;
-- CALL testproc("20210630163227149563000042432429");

*/

//TestProc 14.测试调用存储过程
func TestProc(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	demo := &models.DemoStruct{}
	finder := zorm.NewFinder().Append("call testproc(?) ", "20210630163227149563000042432429")
	zorm.QueryRow(ctx, finder, demo)
	fmt.Println(demo)
}
