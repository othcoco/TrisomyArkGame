package main

import (
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	"github.com/xiaonanln/goworld/engine/storage"
	"github.com/xiaonanln/goworld/engine/gwlog"
	//"math/rand"
	//"github.com/json-iterator/go"
	"github.com/xiaonanln/goworld/newGame2/model"
	"github.com/xiaonanln/goworld/newGame2/config"
	"github.com/xiaonanln/goworld/newGame2/tools"
	//"strconv"
	"math"
	"math/rand"
	"time"
	//"reflect"

)


// Role 对象代表一名角色
type Role struct {
	BasicEntity
	RoleBasicData *[]model.RoleAttrModel
	FightID common.EntityID
}

func (this *Role) DescribeEntityType(desc *entity.EntityTypeDesc) {
	desc.SetPersistent(true)
	desc.SetUseAOI(true, 100)
	desc.DefineAttr("BID", "Client", "Persistent") //基础ID
	desc.DefineAttr("playerID", "Persistent") //
	desc.DefineAttr("Status", "Client", "Persistent") //
	desc.DefineAttr("Bound", "Client", "Persistent") //
	desc.DefineAttr("LV", "Client", "Persistent")//
	desc.DefineAttr("EXP", "Client", "Persistent")
	desc.DefineAttr("NiceName", "Client", "Persistent")
	desc.DefineAttr("AddAttrs", "Persistent")
	desc.DefineAttr("Armings", "Persistent")
	desc.DefineAttr("XD", "Client", "Persistent")//盾

	desc.DefineAttr("Color", "Client", "Persistent")
	desc.DefineAttr("Star", "Client", "Persistent")

	
	desc.DefineAttr("ArmingsAttrs")
	desc.DefineAttr("QX", "Client")	
	desc.DefineAttr("FZ", "Client")//重
	desc.DefineAttr("ZBZL", "Client", "Persistent")//装备重量
	desc.DefineAttr("SD", "Client")//速度
	desc.DefineAttr("MZ", "Client")
	

	desc.DefineAttr("WG", "Client")
	desc.DefineAttr("DG", "Client")//电
	desc.DefineAttr("SG", "Client")//水
	desc.DefineAttr("HG", "Client")//火

	desc.DefineAttr("WF", "Client")
	desc.DefineAttr("DF", "Client")
	desc.DefineAttr("SF", "Client")
	desc.DefineAttr("HF", "Client")
	

	desc.DefineAttr("WC", "Client")//物穿
	desc.DefineAttr("BL", "Client")//暴率
	desc.DefineAttr("BS", "Client")//暴伤
	desc.DefineAttr("Skills", "Client")
	desc.DefineAttr("WeaponUsedTime")

	desc.DefineAttr("spaceKind")
		
}

// OnCreated 在Role对象创建后被调用
func (this *Role) OnCreated() {
	this.BasicEntity.OnCreated()
	this.setDefaultAttrs()
	this.SetClientSyncing(true)
}

func (this *Role) OnDestroy(){
	this.BasicEntity.OnDestroy()
}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *Role) setDefaultAttrs() {
	this.Attrs.SetDefaultInt("BID", 1)
	this.Attrs.SetDefaultInt("Status", 0)
	this.Attrs.SetDefaultInt("Bound", 1)
	this.Attrs.SetDefaultInt("LV", 1)
	this.Attrs.SetDefaultInt("EXP", 0)
	this.Attrs.SetDefaultStr("NiceName", "")
	this.Attrs.SetDefaultMapAttr("AddAttrs", goworld.MapAttr())
	this.Attrs.SetDefaultListAttr("Armings", tools.Array(
		0, //0占位
		goworld.MapAttr(), goworld.MapAttr(), goworld.MapAttr(), goworld.MapAttr(), goworld.MapAttr(), goworld.MapAttr()),//6类装备
	)
	this.Attrs.SetDefaultInt("XD", 0)

	this.Attrs.SetDefaultListAttr("ArmingsAttrs", tools.Array(
		0, //0属性未初始化 1已经初始化
		goworld.MapAttr(),goworld.MapAttr(),goworld.MapAttr(),goworld.MapAttr(),goworld.MapAttr(),goworld.MapAttr(),//6类装备attr
	))
	this.Attrs.SetDefaultInt("Color", 0)
	this.Attrs.SetDefaultInt("Star", 1)

	this.Attrs.SetDefaultInt("QX", 0)
	this.Attrs.SetDefaultInt("FZ", 0)
	this.Attrs.SetDefaultInt("ZBZL", 0)
	this.Attrs.SetDefaultInt("SD", 0)

	this.Attrs.SetDefaultInt("WG", 0)
	this.Attrs.SetDefaultInt("DG", 0)
	this.Attrs.SetDefaultInt("SG", 0)
	this.Attrs.SetDefaultInt("HG", 0)

	this.Attrs.SetDefaultInt("WF", 0)
	this.Attrs.SetDefaultInt("DF", 0)
	this.Attrs.SetDefaultInt("SF", 0)
	this.Attrs.SetDefaultInt("HF", 0)

	this.Attrs.SetDefaultFloat("WC", 0)
	this.Attrs.SetDefaultFloat("BL", 0)
	this.Attrs.SetDefaultFloat("BS", 0)
	this.Attrs.SetDefaultListAttr("Skills", goworld.ListAttr())
	skills := this.Attrs.GetMapAttr("AddAttrs").GetListAttr("Skills").ToList()
	for _,v := range skills{
		this.Attrs.GetListAttr("Skills").AppendInt(v.(int64))	
	}
	this.Attrs.SetDefaultListAttr("WeaponUsedTime", goworld.ListAttr())
	this.Attrs.SetDefaultInt("spaceKind", 1)
	
	
}


func (this *Role) RelaodFinalAttrs(){

	//gwlog.Debugf("1111111111111 %+v", this.BasicDataRoleAttrs)
	roleBasicAttr := config.BASIC_DATA_ROLE_ATTRS[this.Attrs.GetInt("BID")]
	//gwlog.Debugf("1111111111111 %+v", roleBasicAttr)
	per := config.ROLE_LV_PER * float64(this.Attrs.GetInt("LV"))+1.0
	per = per*math.Pow(config.STAR_ATTR_POW+1.0, float64(this.Attrs.GetInt("Star")))
	//通过等级计算最终属性
	this.Attrs.SetInt("QX", int64(float64(roleBasicAttr.QX + this.Attrs.GetMapAttr("AddAttrs").GetInt("QX")) * per))
	this.Attrs.SetInt("FZ", int64(float64(roleBasicAttr.FZ + this.Attrs.GetMapAttr("AddAttrs").GetInt("FZ")) * per))
	this.Attrs.SetInt("SD", int64(float64(roleBasicAttr.SD + this.Attrs.GetMapAttr("AddAttrs").GetInt("SD")) * per))
	this.Attrs.SetInt("MZ", int64(float64(roleBasicAttr.MZ + this.Attrs.GetMapAttr("AddAttrs").GetInt("MZ")) * per))

	this.Attrs.SetInt("WG", int64(float64(roleBasicAttr.WG + this.Attrs.GetMapAttr("AddAttrs").GetInt("WG")) * per))
	this.Attrs.SetInt("DG", int64(float64(roleBasicAttr.DG + this.Attrs.GetMapAttr("AddAttrs").GetInt("DG")) ))
	this.Attrs.SetInt("SG", int64(float64(roleBasicAttr.SG + this.Attrs.GetMapAttr("AddAttrs").GetInt("SG")) ))
	this.Attrs.SetInt("HG", int64(float64(roleBasicAttr.HG + this.Attrs.GetMapAttr("AddAttrs").GetInt("HG")) ))

	this.Attrs.SetInt("WF", int64(float64(roleBasicAttr.WF + this.Attrs.GetMapAttr("AddAttrs").GetInt("WF")) * per))
	this.Attrs.SetInt("DF", int64(float64(roleBasicAttr.DF + this.Attrs.GetMapAttr("AddAttrs").GetInt("DF")) ))
	this.Attrs.SetInt("SF", int64(float64(roleBasicAttr.SF + this.Attrs.GetMapAttr("AddAttrs").GetInt("SF")) ))
	this.Attrs.SetInt("HF", int64(float64(roleBasicAttr.HF + this.Attrs.GetMapAttr("AddAttrs").GetInt("HF")) ))
	//合并属性，不成长
	this.Attrs.SetFloat("WC", roleBasicAttr.WC + this.Attrs.GetMapAttr("AddAttrs").GetFloat("WC"))
	this.Attrs.SetFloat("BL", roleBasicAttr.BL + this.Attrs.GetMapAttr("AddAttrs").GetFloat("BL"))
	this.Attrs.SetFloat("BS", roleBasicAttr.BS + this.Attrs.GetMapAttr("AddAttrs").GetFloat("BS"))
	for _,v := range roleBasicAttr.Skills{
		this.Attrs.GetListAttr("Skills").AppendInt(v)
	}

	//this.Attrs.GetListAttr("Skills").AppendListAttr(roleBasicAttr.Skills)
	//this.Attrs.SetInt("Skills", roleBasicAttr.Skills.AppendListAttr(this.Attrs.GetMapAttr("AddAttrs").GetListAttr("Skills")))

	gwlog.Debugf("99999999 %+v", this.Attrs)
	this.LoadRoleArmingAttr(true)


	
	/*for _i, _item := range RD {
		//item := _item.([]interface{})
		gwlog.Debugf("ggggggggg i %s : %s", _i, _item)
	}*/
}

func (this *Role) LoadRoleArmingAttr(isReload bool){
	Armings := this.Attrs.GetListAttr("Armings")
	ArmingsAttrs := this.Attrs.GetListAttr("ArmingsAttrs")
	if ArmingsAttrs.GetInt(0) != 0 && isReload != true { return }
	//test
	Armings.GetMapAttr(1).SetStr("ID", "X406Xctt9xZIAAAF")
	Armings.GetMapAttr(1).SetInt("BID", 1)
	Armings.GetMapAttr(1).SetInt("Color", 2)
	Armings.GetMapAttr(1).SetInt("Bound", 1)
	Armings.GetMapAttr(3).SetStr("ID", "X41AaMtt9xpZAAAD")
	Armings.GetMapAttr(3).SetInt("BID", 1)
	Armings.GetMapAttr(3).SetInt("Color", 2)
	Armings.GetMapAttr(3).SetInt("Bound", 1)
	//test end
	for k:=1; k<Armings.Size(); k++{
		if Armings.GetMapAttr(k).Size() > 0 {
			this.LaodArmingToMapByID(Armings.GetMapAttr(k).GetStr("ID"), k)
		}
	}
	
}

func (this *Role) LaodArmingToMapByID(armID string, armKey int) {
	armIDobj := common.EntityID(armID)
	storage.Load("Goods", armIDobj, func(_data interface{}, err error) {
		if err != nil{
			gwlog.Debugf("LaodArmingToMapByID error: %s %s", armID, err)
			return
		}
		//gwlog.Debugf("7777777777777777 %s %s", _data, armIDobj)
		data := _data.(map[string]interface{}) 
		basicid := int(data["basicid"].(int64))
		attr := this.Attrs.GetListAttr("ArmingsAttrs").GetMapAttr(armKey)
		attr.AssignMap(config.BASIC_DATA_GOODS_ATTRS[basicid].FightAttr)
		if armKey > 3{//装备4,5,6在此增加属性，武器1,2,3，在攻击时算属性
			attr.ForEach(func(key string, val interface{}){
				tools.ObjectJoinObjectByKey(this.Attrs, val, key)
			})
		}
		//gwlog.Debugf("6666666666666 %+v", this.Attrs.GetListAttr("ArmingsAttrs"))
	})
	
}


func (this *Role) InitRole(agrs map[string] interface{}){

	for kk,vv := range agrs{
		switch v := vv.(type) {

	        case int:
	                this.Attrs.SetInt(kk, int64(v))
	 
	        case string:
	                this.Attrs.SetStr(kk, v)

	 		case int64:
	                this.Attrs.SetInt(kk, v)
	                
	        default:
	                gwlog.Debugf("InitRole unknown type: %s value: %s", kk, v)
        }
	}
	this.AddAttrsByColor()
}


func (this *Role) AddAttrsByColor (){


	itemMaxPer := 0.85
	color := this.Attrs.GetInt("Color")
	if color < 1 { return }
	per := itemMaxPer*100*float64(color)//5种颜色，每种颜色加%几的属性 90代表颜色最多加到9成
	rand.Seed(time.Now().Unix())
	s := rand.Intn(len(config.RAND_ATTR))
	AddAttrs := this.Attrs.GetMapAttr("AddAttrs")
	if s == 0 { s = 1 }
	role := config.BASIC_DATA_ROLE_ATTRS[this.Attrs.GetInt("BID")]
	isdo := 0
	for i:=0; i<s; i++{
		choiceAttr := config.RAND_ATTR[rand.Intn(len(config.RAND_ATTR))]
		var value float64
		if (i+1) == s{
			value = per
		}else{
			value = float64(rand.Intn(int(per)))
			per -= value
		}
		var result int
		result = tools.Contains(model.FloatItem, choiceAttr)
		if result > -1{
			v, err := tools.GetStructFloat64Field(role, choiceAttr)
			if err != nil || v == 0.0{
				gwlog.Debugf("AddAttrsByColorDo 1: %s v: %s", err, v)
				per += value
				continue
			}
			newv := per/100.0 * v + AddAttrs.GetFloat(choiceAttr)
			if newv > v*itemMaxPer { newv = v*itemMaxPer }
			AddAttrs.SetFloat(choiceAttr, newv)
			isdo = 1
		}else{
			v, err := tools.GetStructInt64Field(role, choiceAttr)
			if err != nil || v == 0{
				gwlog.Debugf("AddAttrsByColorDo 2: %s v: %s", err, v)
				per += value
				continue
			}
			newv := int64(per/100.0 * float64(v)) + AddAttrs.GetInt(choiceAttr)
			if newv > int64(float64(v)*itemMaxPer) { newv = int64(float64(v)*itemMaxPer) }
			AddAttrs.SetInt(choiceAttr, newv)
			isdo = 1
		}
		//this.Attrs.GetMapAttr("AddAttrs").
	}
	if isdo == 0{
		v,_ := tools.GetStructInt64Field(role, "WG")
		AddAttrs.SetInt("WG", int64(float64(v)*per/100.0))
	}
	gwlog.Debugf("TTTTTTTTTT %s %s AddAttrs:%+v", s, per, AddAttrs)
	this.RelaodFinalAttrs()
}

func (this *Role) Attack_Client(weaponID int, targetID common.EntityID) {
	now := time.Now().UnixNano() / 1e6
	if now < (this.GetListAttr("ArmingsAttrs").GetMapAttr(weaponID).GetInt("LQ")+this.GetListAttr("WeaponUsedTime").GetInt(weaponID)){
		return
	}
	this.Attrs.GetListAttr("WeaponUsedTime").SetInt(weaponID, now)
	
	//this.CallAllClients("Shoot")
	this.Call(this.FightID, "Attack", this.ID, weaponID, targetID)

}
