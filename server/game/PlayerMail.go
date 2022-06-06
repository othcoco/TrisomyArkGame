package main

import (
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	//"github.com/xiaonanln/goworld/engine/gwlog"
	//"math/rand"
	//"github.com/json-iterator/go"
	//"github.com/xiaonanln/goworld/newGame2/model"
	//"github.com/xiaonanln/goworld/newGame2/config"
	"github.com/xiaonanln/goworld/newGame2/tools"
	//"strconv"
	//"math"
	//"math/rand"
	"time"
	//"reflect"

)


// PlayerMail 
type PlayerMail struct {
	BasicEntity
	
}

func (this *PlayerMail) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetPersistent(true)
	desc.DefineAttr("items", "Client", "Persistent") 
		
}

// OnCreated 在PlayerMail对象创建后被调用
func (this *PlayerMail) OnCreated() {
	this.Entity.OnCreated()
	this.setDefaultAttrs()
	this.SetClientSyncing(true)
}

func (this *PlayerMail) OnDestroy(){

}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *PlayerMail) setDefaultAttrs() {
	this.Attrs.SetDefaultListAttr("items", goworld.ListAttr())
	
	
}

func (this *PlayerMail) newPlayerMail(title string, senderName string, mID string) {
	item := goworld.MapAttr()
	item.SetStr("title", title)
	item.SetStr("sender", senderName)
	item.SetStr("mid", mID)
	item.SetInt("ctime", int64(time.Now().Unix()))
	item.SetInt("status", 0)
	this.Attrs.GetListAttr("items").AppendMapAttr(item)
	this.Save()
}

/**
*@agrs 
**/
func (this *PlayerMail) NewMail(title string, content string, senderName string, agrs map[string] interface{}) common.EntityID{
	mail := goworld.CreateEntityLocally("MailContent")
	this.newPlayerMail(title, senderName, string(mail.ID))
	this.Call(mail.ID, "InitMailContent", content, agrs)
	return mail.ID
}



func (this *PlayerMail) InitPlayerMail(playerID common.EntityID, playerMailID common.EntityID, spaceKind int64){
		//this.Attrs.SetInt("spaceKind", spaceKind)
	this.CreateClient(playerID)
	//this.EnterSpace_Client(int(spaceKind))
	this.SetClientSyncing(true)
	//gwlog.Infof("bbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbbb: %s", this.GetClient())
	this.NewMail("test_title", "hello content", "system", map[string] interface{} {
		"gids": tools.Array(tools.Array("X4AMf8tt9w98AAAD",1,2)),
	})
}

//goodsID 为 "X4AMf8tt9w98AAAD" 已创建好物品， 为数字时，物品收取时创建此物品
func (this *PlayerMail) NewGoodsFullMail(goodsID interface{}, goodsBasicID int, goodsNum int){
	this.NewMail("title_goods_full", "content_goods_full", "sender_system", map[string] interface{} {
		"gids": tools.Array(tools.Array(goodsID, goodsBasicID, goodsNum)),
	})
}