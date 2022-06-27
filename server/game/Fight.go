package main

import (
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	"github.com/xiaonanln/goworld/engine/gwlog"
	//"math/rand"
	//"github.com/json-iterator/go"
	//"github.com/xiaonanln/goworld/newGame2/model"
	"github.com/xiaonanln/goworld/newGame2/config"
	//"github.com/xiaonanln/goworld/newGame2/tools"
	//"strconv"
	//"math"
	"math/rand"
	"time"
	//"reflect"

)


// Fight 
type Fight struct {
	BasicEntity
	MID int64
	playerID common.EntityID
}

func (this *Fight) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetPersistent(false)
	desc.SetUseAOI(true, 255*255*255*255)
	desc.DefineAttr("t1", "AllClients")//
	desc.DefineAttr("t2", "AllClients")//
	desc.DefineAttr("ctime", "AllClients")//
	
		
}

// OnCreated 在Fight对象创建后被调用
func (this *Fight) OnCreated() {
	rand.Seed(time.Now().Unix())
	this.BasicEntity.OnCreated()
	this.setDefaultAttrs()
	this.SetClientSyncing(true)
}

func (this *Fight) OnDestroy(){
	this.BasicEntity.OnDestroy()
}

// setDefaultAttrs 
func (this *Fight) setDefaultAttrs() {
	this.Attrs.SetDefaultListAttr("t1", goworld.ListAttr())
	this.Attrs.SetDefaultListAttr("t2", goworld.ListAttr())
	this.Attrs.SetInt("ctime", time.Now().Unix())
	
	
}

func (this *Fight) Attack (roleID common.EntityID, weaponID int, targetID common.EntityID){
	role := goworld.GetEntity(roleID)
	target := goworld.GetEntity(targetID)
	if target == nil{ 
		role.CallClient("miss_target") 
		return
	}
	if target.Attrs.GetInt("QX") < 1 { return }
	this.roleAttackRole(role, weaponID, target)

}



func (this *Fight) roleAttackRole (role *entity.Entity, weaponID int, target *entity.Entity){
	mz := 1.0 - float32(target.Attrs.GetInt("SD")/role.Attrs.GetInt("MZ"))
	if mz < 0.5{ mz = 0.5}
	if mz < rand.Float32(){ //未命中
		target.CallClient("attack_miss")
		return
	}
	wg := role.GetInt("WG")+role.GetListAttr("ArmingsAttrs").GetMapAttr(weaponID).GetInt("WG")
	dg := role.GetInt("DG")+role.GetListAttr("ArmingsAttrs").GetMapAttr(weaponID).GetInt("DG")
	sg := role.GetInt("SG")+role.GetListAttr("ArmingsAttrs").GetMapAttr(weaponID).GetInt("SG")
	hg := role.GetInt("HG")+role.GetListAttr("ArmingsAttrs").GetMapAttr(weaponID).GetInt("HG")
	hurt := this.hurtCommon(wg, target.GetInt("WF"), target.GetFloat("WC"))
	hurt += this.hurtCommon(dg, target.GetInt("DF"), 0.1)
	hurt += this.hurtCommon(sg, target.GetInt("SF"), 0.1)
	hurt += this.hurtCommon(hg, target.GetInt("hF"), 0.1)
	hurt = this.hurtBJ(hurt, role.Attrs.GetFloat("BL"), role.Attrs.GetFloat("BS"))
	xd := target.GetInt("XD")
	if xd > 0{
		target.Attrs.SetInt("XD", xd-hurt)
	}else{
		target.Attrs.SetInt("QX", target.Attrs.GetInt("QX")-hurt)
	}
	target.SetClientSyncing(true)
}

//物理攻击结果计算
// g攻击 f防御 c穿透
func (this *Fight) hurtCommon(g int64, f int64, c float64) int64{
	if g == 0 { return 0}
	hurt := g - f
	if hurt < 0 { hurt = 0}
	hurt += int64(float64(hurt)*c)
	return hurt
}

//暴击计算
func (this *Fight) hurtBJ(hurt int64, bl float64, bs float64) int64{
	if bl < rand.Float64(){
		if bs > 2.5 { bs = 2.5 + (bs-2.5)/5.0 }
		return hurt+int64(float64(hurt)*bs)
	}
	return hurt
}

func (this *Fight) createFightSpace(playerID common.EntityID) *entity.Space{
	kind := rand.Intn(255*255*255*255*255*255*255*128)
	space := goworld.CreateSpaceLocally(kind)
	space.EnableAOI(255*255*255*255)
	this.EnterSpace(space.ID, entity.Vector3{})
	this.CreateClient(playerID)	
	return space
}

func (this *Fight) StartFightByMID(monsterList []int64, roleList *entity.ListAttr, MID int64, playerID common.EntityID){
	space := this.createFightSpace(playerID)
	this.MID = MID
	this.playerID = playerID
	
	for i := 0; i<roleList.Size(); i++{
		role := goworld.GetEntity(common.EntityID(roleList.GetStr(i)))
		r := role.I.(*Role)
		if r.FightID == common.EntityID(""){
			r.EnterSpace(space.ID, entity.Vector3{})
			r.FightID = this.ID
			this.Attrs.GetListAttr("t1").AppendStr(roleList.GetStr(i))
			//gwlog.Infof("%s set FightID:%s", role, this.ID)
		}else{
			gwlog.Infof("error: role %s is fighting! FightID:%s this.ID:%s", role, r.FightID, this.ID)
		}
	}
	
	//创建怪物
	for _,v := range monsterList{
		eid := common.GenEntityID()
		m := this.Space.CreateEntityIntact("Monsters", entity.Vector3{}, eid, map[string]interface{}{
			"MID": MID,
			"BID": v,
			"AddAttrs": config.BASIC_DATA_MONSTERS[v].AddAttrs,
			"Armings": config.BASIC_DATA_MONSTERS[v].Armings,
			"Goods": config.BASIC_DATA_MONSTERS[v].Goods,
			"Roles": config.BASIC_DATA_MONSTERS[v].Roles,
			"Coin1": config.BASIC_DATA_MONSTERS[v].Coin1,
			"Coin2": config.BASIC_DATA_MONSTERS[v].Coin2,
			"Coin3": config.BASIC_DATA_MONSTERS[v].Coin3,
		})
		this.Attrs.GetListAttr("t2").AppendStr(string(eid))
		monster := m.I.(*Monsters)
		monster.FightID = this.ID
		monster.InitMonster(map[string]interface{}{})
		//monster.CreateClient(this.playerID)
		monster.EnterSpace(space.ID, entity.Vector3{})
		
	}

	gwlog.Infof("fight: %s start", this.ID)

	
}