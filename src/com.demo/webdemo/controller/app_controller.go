package controller

import (
	"com.demo/webdemo/models"
	"com.demo/webdemo/service"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type AppController struct {
	web.Controller
}

func (this AppController) ListAll() {
	appService := new(service.AppService)
	appList := appService.ListAll()

	logs.Info("查询集合 - ", appList)

	r := new(models.R)
	r.Data = appList

	this.Ctx.JSONResp(r)
}
