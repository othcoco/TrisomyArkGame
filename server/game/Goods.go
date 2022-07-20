package main

import (
	"github.com/xiaonanln/goworld"
	//"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	"github.com/xiaonanln/goworld/engine/gwlog"
	//"github.com/xiaonanln/goworld/newGame2/config"
	//"strconv"
)

// Player 对象代表一名玩家
type Goods struct {
	entity.Entity
	attr *entity.MapAttr
}

func (this *Goods) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetPersistent(true)
	desc.DefineAttr("basicid", "Persistent")//物品BasicID
	desc.DefineAttr("Bound", "Persistent")//绑定
	desc.DefineAttr("gstatus", "Persistent")//物品使用状态
	desc.DefineAttr("gattrs", "Persistent")//物品属性
	desc.DefineAttr("playerID", "Persistent")//所属用户
	desc.DefineAttr("ctime", "Persistent")//创建时间


}

// OnCreated 在Player对象创建后被调用
func (this *Goods) OnCreated() {
	this.Entity.OnCreated()
	this.setDefaultAttrs()
}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *Goods) setDefaultAttrs() {
	this.Attrs.SetDefaultInt("basicid", 1)
	this.Attrs.SetDefaultInt("Bound", 1)
	this.Attrs.SetDefaultInt("gstatus", 0)
	this.Attrs.SetDefaultListAttr("gattrs", goworld.ListAttr())
	this.Attrs.SetDefaultStr("playerID", "")
	this.Attrs.SetDefaultInt("ctime", 1000)
	//this.SetClientSyncing(true)
}

func (this *Goods) InitGoods(agrs map[string] interface{}){

	for kk,vv := range agrs{
		switch v := vv.(type) {

	        case int:
	                this.Attrs.SetInt(kk, int64(v))
	 
	        case string:
	                this.Attrs.SetStr(kk, v)

	        case int64:
	                this.Attrs.SetInt(kk, v)	 
	        default:
	               gwlog.Debugf("InitGoods unknown type: %s value: %s", kk, v)
        }
	}
	this.Save()
	gwlog.Debugf("InitGoods >>>>>>> %s", this.Attrs)
}

/*func (this *Goods) GerFinalArmingAtrr (armKey int, armID common.EntityID, callBack string) {
	//arming := config.BASIC_DATA_GOODS_ATTRS[this.Attrs.GetInt("basicid")]
	gwlog.Debugf("888888888888888888888888888888 %+v", this.Attrs.GetInt("basicid"))
	gwlog.Debugf("888888888888888888888888888888 %+v", this.GetInt("basicid"))
	gwlog.Debugf("888888888888888888888888888888 %+v", config.BASIC_DATA_GOODS_ATTRS[armKey])
	attr := goworld.MapAttr()
	this.Call(armID, callBack, armKey, attr)
}*/