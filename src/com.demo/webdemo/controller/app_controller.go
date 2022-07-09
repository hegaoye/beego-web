package controller

import (
	"com.demo/webdemo/models"
	"com.demo/webdemo/service"
	"github.com/beego/beego/core/logs"
	"github.com/beego/beego/server/web"
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
