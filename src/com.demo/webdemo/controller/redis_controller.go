package controller

import (
	"com.demo/webdemo/models"
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/core/logs"
	"time"
)

type RedisController struct {
	beego.Controller
}

func (this RedisController) Get() {
	key := this.Ctx.Input.Param(":key")
	value := models.GetKey(key)

	logs.Info("get key - ", value)

	this.Ctx.JSONResp(models.R{Data: value})
}

func (this RedisController) Add() {
	key := this.Ctx.Input.Param(":key")
	value := this.Ctx.Input.Param(":value")

	logs.Info("add key - ", key, ",value - ", value)

	models.SetStr(key, value, time.Duration(0))
	mapStr := map[string]string{}
	mapStr[key] = value

	this.Ctx.JSONResp(models.R{Data: mapStr})
}

func (this RedisController) Del() {
	key := this.Ctx.Input.Param(":key")
	logs.Info("del key - ", key)

	models.DelKey(key)

	this.Ctx.JSONResp(models.R{Data: key})
}
