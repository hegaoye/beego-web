package controller

import (
	"com.demo/webdemo/models"
	"com.demo/webdemo/rpcclient"
	"github.com/beego/beego/v2/server/web"
)

type GrpcTestController struct {
	web.Controller
}

func (this GrpcTestController) SendMsg() {
	data := this.Ctx.Input.Param(":data")
	goBaseResult := rpcclient.SendToServerMessageResult(data)
	r := new(models.R)
	r.Code = 200
	r.Data = goBaseResult

	this.Ctx.JSONResp(r)
}
