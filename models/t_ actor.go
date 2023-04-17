package models

import (
	"time"

	"gitee.com/chunanyong/zorm"
)

const ActorStructTableName = "t_actor"

type ActorStruct struct {
	zorm.EntityStruct
	Id         string    `column:"id"`
	StageName  string    `column:"stageName"`
	RealName   string    `column:"realName"`
	Company    string    `column:"company"`
	CreateTime time.Time `column:"createTime"`
}

func (entity *ActorStruct) GetTableName() string {
	return ActorStructTableName
}

func (entity *ActorStruct) GetPKColumnName() string {
	return "id"
}

type ActorInfo struct {
	// ActorStruct
	Id         string    `column:"id"`
	StageName  string    `column:"stageName"`
	RealName   string    `column:"realName"`
	Company    string    `column:"company"`
	CreateTime time.Time `column:"createTime"`
	// DemoStruct
	DemoStructId         string    `column:"id"`
	DemoStructUserName   string    `column:"userName"`
	DemoStructPassword   string    `column:"password"`
	DemoStructCreateTime time.Time `column:"createTime"`
	DemoStructActive     int
}
