package controller

import (
	"github.com/beego/beego/v2/adapter/logs"
	//"github.com/beego/beego/core/logs"
	//"github.com/beego/beego/server/web"
	"github.com/beego/beego/v2/server/web"
	"time"
)

type HeartbeatController struct {
	web.Controller
}

// 心跳
func (this HeartbeatController) Heartbeat() {
	logs.Info("心跳时间 - ", time.Now().Local())
	this.Ctx.WriteString("OK")
}
