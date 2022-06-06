package main

import (
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	"github.com/xiaonanln/goworld/engine/gwlog"
	"math/rand"
	//"strconv"
)

// PlayerAttr 对象代表一名玩家
type PlayerAttr struct {
	entity.Entity
}

func (this *PlayerAttr) DescribeEntityType(desc *entity.EntityTypeDesc) {
	//desc.SetPersistent(true)
	desc.DefineAttr("qx", "Client", "Persistent")
	desc.DefineAttr("xd", "Client", "Persistent")//烯盾
	desc.DefineAttr("fz", "Client", "Persistent")//负重

	desc.DefineAttr("wg", "Client", "Persistent")
	desc.DefineAttr("dg", "Client", "Persistent")//电
	desc.DefineAttr("sg", "Client", "Persistent")//水
	desc.DefineAttr("hg", "Client", "Persistent")//火

	desc.DefineAttr("wf", "Client", "Persistent")
	desc.DefineAttr("df", "Client", "Persistent")
	desc.DefineAttr("sf", "Client", "Persistent")
	desc.DefineAttr("hf", "Client", "Persistent")
	desc.DefineAttr("sd", "Client", "Persistent")//速度

	desc.DefineAttr("wc", "Client", "Persistent")//物穿
	desc.DefineAttr("bl", "Client", "Persistent")//暴率
	desc.DefineAttr("bs", "Client", "Persistent")//暴伤
		
}

// OnCreated 在PlayerAttr对象创建后被调用
func (this *PlayerAttr) OnCreated() {
	this.Entity.OnCreated()
	this.setDefaultAttrs()
}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *PlayerAttr) setDefaultAttrs() {
	this.Attrs.SetDefaultInt("qx", 0)
	this.Attrs.SetDefaultInt("xd", 0)
	this.Attrs.SetDefaultInt("fz", 0)

	this.Attrs.SetDefaultInt("wg", 0)
	this.Attrs.SetDefaultInt("dg", 0)
	this.Attrs.SetDefaultInt("sg", 0)
	this.Attrs.SetDefaultInt("hg", 0)

	this.Attrs.SetDefaultInt("wf", 0)
	this.Attrs.SetDefaultInt("df", 0)
	this.Attrs.SetDefaultInt("sf", 0)
	this.Attrs.SetDefaultInt("hf", 0)
	this.Attrs.SetDefaultInt("sd", 0)

	this.Attrs.SetDefaultFloat("wc", 0)
	this.Attrs.SetDefaultFloat("bl", 0)
	this.Attrs.SetDefaultFloat("bs", 0)
	this.SetClientSyncing(true)
}

