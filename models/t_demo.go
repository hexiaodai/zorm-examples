package models

import (
	"time"

	"gitee.com/chunanyong/zorm"
)

//demoStructTableName 表名常量,方便直接调用
const DemoStructTableName = "t_demo"

// demoStruct 例子
type DemoStruct struct {
	//引入默认的struct,隔离IEntityStruct的方法改动
	zorm.EntityStruct

	//Id 主键
	Id string `column:"id"`

	//UserName 姓名
	UserName string `column:"userName"`

	//Password 密码
	Password string `column:"password"`

	//CreateTime <no value>
	CreateTime time.Time `column:"createTime"`

	//Active 是否有效(0否,1是)
	//Active int `column:"active"`

	//------------------数据库字段结束,自定义字段写在下面---------------//
	//如果查询的字段在column tag中没有找到,就会根据名称(不区分大小写,支持 _ 下划线转驼峰)映射到struct的属性上

	//模拟自定义的字段Active
	Active int `column:"active"`
}

//GetTableName 获取表名称
//IEntityStruct 接口的方法,实体类需要实现!!!
func (entity *DemoStruct) GetTableName() string {
	return DemoStructTableName
}

//GetPKColumnName 获取数据库表的主键字段名称.因为要兼容Map,只能是数据库的字段名称
//不支持联合主键,变通认为无主键,业务控制实现(艰难取舍)
//如果没有主键,也需要实现这个方法, return "" 即可
//IEntityStruct 接口的方法,实体类需要实现!!!
func (entity *DemoStruct) GetPKColumnName() string {
	//如果没有主键
	//return ""
	return "id"
}
