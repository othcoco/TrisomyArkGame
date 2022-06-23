package main

import (
	//"github.com/xiaonanln/goworld"
	//"github.com/xiaonanln/goworld/engine/common"
	//"github.com/xiaonanln/goworld/engine/consts"
	"github.com/xiaonanln/goworld/engine/entity"
	//"github.com/xiaonanln/goworld/engine/gwlog"
	//"math/rand"
	//"github.com/json-iterator/go"
	//"github.com/xiaonanln/goworld/newGame2/model"
	//"github.com/xiaonanln/goworld/newGame2/config"
	//"github.com/xiaonanln/goworld/newGame2/tools"
	//"strconv"
	//"math"
	//"math/rand"
	//"time"
	//"reflect"

)


// Monsters 
type Monsters struct {
	Role
	
}

func (this *Monsters) DescribeEntityType(desc *entity.EntityTypeDesc) {
	this.Role.DescribeEntityType(desc)
	desc.SetPersistent(false)	
		
}

// OnCreated 在Monsters对象创建后被调用
func (this *Monsters) OnCreated() {
	this.Role.OnCreated()
}

func (this *Monsters) OnDestroy(){
	this.Role.OnDestroy()
}

// setDefaultAttrs 设置玩家的一些默认属性
func (this *Monsters) setDefaultAttrs() {
	this.Role.setDefaultAttrs()
	
}

func (this *Monsters) InitMonster(agrs map[string] interface{}){
	this.Role.InitRole(agrs)
}

