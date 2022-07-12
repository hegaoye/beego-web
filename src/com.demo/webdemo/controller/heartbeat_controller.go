package controller

import (
	"github.com/beego/beego/v2/core/logs"
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
