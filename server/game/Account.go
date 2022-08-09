package main

import (
	"github.com/xiaonanln/goworld"
	"github.com/xiaonanln/goworld/engine/common"
	"github.com/xiaonanln/goworld/engine/entity"
	"github.com/xiaonanln/goworld/engine/gwlog"
)

// Account 是账号对象类型，用于处理注册、登录逻辑
type Account struct {
	entity.Entity // 自定义对象类型必须继承entity.Entity
	logIn         bool
}

func (this *Account) DescribeEntityType(desc *entity.EntityTypeDesc) {
}

// Register_Client 是处理玩家注册请求的RPC函数
func (this *Account) Register_Client(username string, password string) {
	gwlog.Debugf("Register %s %s", username, password)
	goworld.GetOrPutKVDB("password$"+username, password, func(oldVal string, err error) {
		if err != nil {
			this.CallClient("ShowError", "Server Error： "+err.Error()) // 服务器错误
			return
		}

		if oldVal == "" {

			player := goworld.CreateEntityLocally("Player") // 创建一个Player对象然后立刻销毁，产生一次存盘
			player.Attrs.SetStr("name", username)
			player.Destroy()

			goworld.PutKVDB("playerID$"+username, string(player.ID), func(err error) {
				this.CallClient("ShowInfo", "Registered Successfully, please click login.") // 注册成功，请点击登录
			})
		} else {
			this.CallClient("ShowError", "Sorry, this account aready exists.") // 抱歉，这个账号已经存在
		}
	})
}

// Login_Client 是处理玩家登录请求的RPC函数
func (this *Account) Login_Client(username string, password string) {
	gwlog.Debugf("%s.Login: username=%s, password=%s", this, username, password)
	if this.logIn {
		// logining
		gwlog.Errorf("%s has already started to log in.", this)
		return
	}

	gwlog.Infof("%s started log in with username %s password %s ...", this, username, password)
	this.logIn = true
	goworld.GetKVDB("password$"+username, func(correctPassword string, err error) {
		if err != nil {
			this.logIn = false
			this.CallClient("ShowError", "Server Error： "+err.Error()) // 服务器错误
			return
		}

		if correctPassword == "" {
			this.logIn = false
			this.CallClient("ShowError", "Account does not exist.") // 账号不存在
			return
		}

		if password != correctPassword {
			this.logIn = false
			this.CallClient("ShowError", "Invalid password or username") // 密码错误
			return
		}

		goworld.GetKVDB("playerID$"+username, func(_playerID string, err error) {
			if err != nil {
				this.logIn = false
				this.CallClient("ShowError", "Server Error："+err.Error()) // 服务器错误
				return
			}
			playerID := common.EntityID(_playerID)
			player := entity.GetEntity(playerID)
			if player != nil{
				this.GiveClientTo(player)
				//gwlog.Debugf("@@@@@@@@@@@@@@@@:  %s ", playerID)
			}else{
				goworld.LoadEntityAnywhere("Player", playerID)
				this.Call(playerID, "GetSpaceID", this.ID)
				//gwlog.Debugf("################:  %s ", playerID)
			}
		})
	})
}

// OnGetPlayerSpaceID 是用于接收Player场景编号的回调函数
func (this *Account) OnGetPlayerSpaceID(playerID common.EntityID, spaceID common.EntityID) {
	// player may be in the same space with account, check again
	player := goworld.GetEntity(playerID)
	if player != nil {
		this.onPlayerEntityFound(player)
		return
	}

	this.Attrs.SetStr("loginPlayerID", string(playerID))
	this.EnterSpace(spaceID, entity.Vector3{})
}

func (this *Account) onPlayerEntityFound(player *entity.Entity) {
	gwlog.Infof("Player %s is found, giving client to ...", player)
	this.logIn = false
	this.GiveClientTo(player) // 将Account的客户端移交给Player
}

// OnClientDisconnected 在客户端掉线或者给了Player后触发
func (this *Account) OnClientDisconnected() {
	gwlog.Debugf("destroying %s ...", this)
	this.Destroy()
}

// OnMigrateIn 在账号迁移到目标服务器的时候调用
func (this *Account) OnMigrateIn() {
	loginPlayerID := common.EntityID(this.Attrs.GetStr("loginPlayerID"))
	player := goworld.GetEntity(loginPlayerID)
	gwlog.Debugf("%s migrating in, attrs=%v, loginPlayerID=%s, player=%v, client=%s", this, this.Attrs.ToMap(), loginPlayerID, player, this.GetClient())

	if player != nil {
		this.onPlayerEntityFound(player)
	} else {
		// failed
		this.CallClient("ShowError", "登录失败，请重试")
		this.logIn = false
	}
}
