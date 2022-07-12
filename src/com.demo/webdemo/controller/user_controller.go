package controller

import (
	"com.demo/webdemo/models"
	"com.demo/webdemo/service"
	"encoding/json"
	"github.com/beego/beego/v2/core/logs"
	"github.com/beego/beego/v2/server/web"
)

type UserController struct {
	web.Controller
}

func (this *UserController) ListAll() {
	userService := new(service.UserService)
	userList := userService.ListAll()

	this.Ctx.JSONResp(userList)
}

func (this *UserController) GetOneByName() {
	name := this.Ctx.Input.Param(":name")

	userService := new(service.UserService)
	user := userService.GetOneByName(name)

	this.Ctx.JSONResp(user)
}

func (this *UserController) GetOneByWorkNo() {
	workNo := this.Ctx.Input.Param(":workNo")

	userService := new(service.UserService)
	user := userService.GetOneByWorkNo(workNo)

	this.Ctx.JSONResp(user)
}

func (this *UserController) UpdateUByName() {
	id := this.Ctx.Input.Param(":name")
	u := this.Ctx.Input.Param(":u")

	userService := new(service.UserService)
	user := userService.UpdateUByName(id, u)

	this.Ctx.JSONResp(user)
}

func (this *UserController) UploadData() {
	var userList []models.Staff
	data := this.Ctx.Input.RequestBody

	err := json.Unmarshal(data, &userList)
	if err != nil {
		logs.Debug("json.Unmarshal is err:", err.Error())
	}

	logs.Info("the upload data - ", userList)

	this.Ctx.WriteString("OK")
}
