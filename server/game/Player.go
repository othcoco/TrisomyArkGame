package main

import (
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	//"github.com/xiaonanln/goworld/engine/gwlog"
	"math/rand"
	//"strconv"

	//"github.com/xiaonanln/goworld/newGame2/model"
)

// Player 对象代表一名玩家
type Player struct {
	BasicEntity
	playerID common.EntityID
	playerBagID common.EntityID
	playerBagRoleID common.EntityID
	playerMailID common.EntityID
	playerMissionID common.EntityID
}

func (this *Player) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetPersistent(true)
	desc.DefineAttr("name", "AllClients", "Persistent")
	desc.DefineAttr("lv", "AllClients", "Persistent")
	desc.DefineAttr("bagID", "Persistent")
	desc.DefineAttr("bagRoleID", "Persistent")
	desc.DefineAttr("fightRoleIDs", "Persistent")
	desc.DefineAttr("missionID", "Persistent")
	desc.DefineAttr("Coin1", "Persistent")
	desc.DefineAttr("Coin2", "Persistent")
	desc.DefineAttr("Coin3", "Persistent")
	desc.DefineAttr("mail", "Persistent")
	desc.DefineAttr("spaceKind")
}

// OnCreated 在Player对象创建后被调用
func (this *Player) OnCreated() {
	this.Entity.OnCreated()
	this.setDefaultAttrs()
	this.playerID = this.ID
}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *Player) setDefaultAttrs() {
	this.Attrs.SetDefaultStr("name", "noname")
	this.Attrs.SetDefaultInt("lv", 1)
	this.Attrs.SetDefaultStr("bagID", "")
	this.Attrs.SetDefaultStr("bagRoleID", "")
	this.Attrs.SetDefaultListAttr("fightRoleIDs", goworld.ListAttr())
	this.Attrs.SetDefaultStr("missionID", "")
	this.Attrs.SetDefaultInt("Coin1", 0)
	this.Attrs.SetDefaultInt("Coin2", 0)
	this.Attrs.SetDefaultInt("Coin3", 0)
	this.Attrs.SetDefaultStr("mail", "")
	//randomNUm := int64(rand.Intn(goworld.GetServiceShardCount("SpaceService")-1)+1)
	randomNUm := int64(rand.Intn(_MAX_AVATAR_COUNT_PER_SPACE))
	this.Attrs.SetDefaultInt("spaceKind", randomNUm)
	this.SetClientSyncing(true)
	//this.Attrs.GetListAttr("fightRoleIDs").AppendStr("X592MMtt9xLvAAAF")
	//this.Attrs.GetListAttr("fightRoleIDs").AppendStr("X594cctt9xUeAAAB")
}



func (this *Player) OnClientConnected() {
	this.BasicEntity.OnClientConnected()
	this.enterSpaceAfter(int(this.GetInt("spaceKind")))
	
}

func (this *Player) enterSpaceAfter(spaceKind int) {
	if this.Space.Kind == spaceKind {
		this.NewAllChildClientToChild()
	}
}

//一般情况为挤号，重新建立客户端
func (this *Player) NewAllChildClientToChild(){
	this.NewClientToChild(this.playerBagID)
	this.NewClientToChild(this.playerBagRoleID)
	this.NewClientToChild(this.playerMailID)
	this.NewClientToChild(this.playerMissionID)
}

func (this *Player) NewClientToChild(childID common.EntityID){
	child := goworld.GetEntity(childID)
	if child != nil {
		child.AssignClient(this.GetClient())
		child.CreateEntityToClient()
	}
}

func (this *Player) OnEnterSpace() {
	this.BasicEntity.OnEnterSpace()
	this.laodPlayerNeedData()
}

func (this *Player) laodPlayerNeedData(){
	this.playerBagID = this.laodNeedEntity("Bag", "bagID", "InitBag")
	this.playerBagRoleID = this.laodNeedEntity("BagRole", "bagRoleID", "InitBagRole")
	this.playerMailID = this.laodNeedEntity("PlayerMail", "mail", "InitPlayerMail")
	this.playerMissionID = this.laodNeedEntity("Mission", "missionID", "InitPlayerMission")
	this.laodFightRole()
}

//物品满时发送邮件
func (this *Player) SendGoodsMaill(goodsID interface{}, goodsBasicID int, goodsNum int){
	this.Call(this.playerMailID, "NewGoodsFullMail", goodsID, goodsBasicID, goodsNum)
}

func (this *Player) laodFightRole (){
	fightRoleIDs := this.Attrs.GetListAttr("fightRoleIDs")
	for i := 0; i<fightRoleIDs.Size(); i++{
		this.Space.CreateEntityIntact("Role", entity.Vector3{}, common.EntityID(fightRoleIDs.GetStr(i)), nil)
	}
}
