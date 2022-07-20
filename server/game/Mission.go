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
	"github.com/xiaonanln/goworld/newGame2/config"
	//"github.com/xiaonanln/goworld/newGame2/tools"
	"strconv"
	"math"
	//"math/rand"
	//"time"
	//"reflect"
	"strings"

)


// Mission 
type Mission struct {
	BasicEntity
	playerID common.EntityID
}

func (this *Mission) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetPersistent(true)
	desc.DefineAttr("playerID", "Client", "Persistent") //基础ID
	desc.DefineAttr("mIDs", "Client", "Persistent")//当前任务
	desc.DefineAttr("doneForeverIDs", "Persistent")//已经完成的不可重复的任务
	desc.DefineAttr("doneDayIDs")//每天可以重复完成的任务

	
		
}

// OnCreated 在Mission对象创建后被调用
func (this *Mission) OnCreated() {
	this.Entity.OnCreated()
	this.setDefaultAttrs()
	this.SetClientSyncing(true)
}

func (this *Mission) OnDestroy(){

}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *Mission) setDefaultAttrs() {
	this.Attrs.SetDefaultStr("playerID", "")
	this.Attrs.SetDefaultMapAttr("mIDs", goworld.MapAttr())
	this.Attrs.SetDefaultMapAttr("doneForeverIDs", goworld.MapAttr())
	this.Attrs.SetDefaultMapAttr("doneDayIDs", goworld.MapAttr())
	
}

func (this *Mission) InitPlayerMission (playerID common.EntityID, playerMailID common.EntityID, spaceKind int64){
	this.CreateClient(playerID)
	this.playerID = playerID
}

func (this *Mission) GetMissionByID(mID int){
	m := config.BASIC_DATA_MISSIONS[mID]
	var keyStr strings.Builder
	keyStr.WriteString("mid_")
	keyStr.WriteString(strconv.Itoa(mID))
	if m.Type > 0{
		nowNum := this.GetMapAttr("doneDayIDs").GetInt(keyStr.String())
		if nowNum < m.Type  && this.GetMapAttr("mIDs").HasKey(keyStr.String()) == false{ 
			this.GetMapAttr("doneDayIDs").SetInt(keyStr.String(), nowNum+1)
			this.GetMapAttr("mIDs").SetInt(keyStr.String(), 0)
			this.Save()
		}else{
			this.CallClient("Server_Msg", "mission_limit_max") 
			return
		}
	}else{
		nowNum := this.GetMapAttr("doneForeverIDs").GetInt(keyStr.String())
		if float64(nowNum) < math.Abs(float64(m.Type)) && this.GetMapAttr("mIDs").HasKey(keyStr.String()) == false{ 
			this.GetMapAttr("doneForeverIDs").SetInt(keyStr.String(), nowNum+1)
			this.GetMapAttr("mIDs").SetInt(keyStr.String(), 0)
			this.Save()
		}else{
			this.CallClient("Server_Msg", "mission_limit_max") 
			return
		}
	}
	
	//gwlog.Infof("MMMMMMMMM %+v", m)
}

func (this *Mission) startMissionFight(monsterList []int64, MID int64){
	player := goworld.GetEntity(this.playerID)
	if player.GetListAttr("fightRoleIDs").Size() < 1{
		this.CallClient("Server_Msg", "fight_role_less")
		return
	}
	if len(monsterList) > 0{
		fightID := common.GenEntityID()
		fight := this.Space.CreateEntityIntact("Fight", entity.Vector3{}, fightID, nil)
		fight.I.(*Fight).StartFightByMID(monsterList, player.GetListAttr("fightRoleIDs"), MID, this.playerID)
		
	}
}

func (this *Mission) GetItem_Client(MID int){
	this.GetMissionByID(MID)
}

func (this *Mission) StartFight_Client(MID int){
	var keyStr strings.Builder
	keyStr.WriteString("mid_")
	keyStr.WriteString(strconv.Itoa(MID))
	if this.Attrs.GetMapAttr("mIDs").HasKey(keyStr.String()) == false { return } //未接受此任务
	m := config.BASIC_DATA_MISSIONS[MID]
	this.startMissionFight(m.Monsters, int64(MID))
}

