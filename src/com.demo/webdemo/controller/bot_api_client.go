package controller

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type BotApiClient struct {
	web.Controller
}

// Get 後去機器人信息
func (receiver BotApiClient) Get() {
	//result := service.BotClient()
	result := ""
	logs.Debug(result)

	receiver.Ctx.WriteString(result)
}
