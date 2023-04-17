package zormexamples

import (
	"context"
	"testing"

	"gitee.com/chunanyong/zorm"
)

/*

-- 自定义函数
set global log_bin_trust_function_creators=1;
DROP FUNCTION IF EXISTS testfunc;

DELIMITER $
CREATE FUNCTION testfunc(id_in VARCHAR(100))
RETURNS VARCHAR(100)
BEGIN
	DECLARE userName_out VARCHAR(100) DEFAULT '20220726125301346422000491406956';
	SELECT userName INTO userName_out FROM t_demo WHERE id=id_in;
	RETURN userName_out;
END$
DELIMITER;

-- 调用函数
-- SELECT testfunc("20220726125301346422000491406956")
*/

//TestFunc 15.测试调用自定义函数
func TestFunc(t *testing.T) {
	// ctx 一般一个请求一个ctx,正常应该有web层传入,例如gin的c.Request.Context().这里只是模拟
	var ctx = context.Background()

	var userName string
	finder := zorm.NewFinder().Append("select testfunc(?) ", "20210630163227149563000042432429")
	has, err := zorm.QueryRow(ctx, finder, &userName)
	t.Logf("has: %v", has)
	t.Logf("err: %v", err)

	t.Log(userName)
}
