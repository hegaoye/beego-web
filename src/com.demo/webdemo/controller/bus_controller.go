package controller

import (
	"com.demo/webdemo/models"
	"com.demo/webdemo/service"
	"encoding/json"
	"github.com/beego/beego/server/web"
)

type BusController struct {
	web.Controller
}

func (this BusController) Get() {
	name := this.Ctx.Input.Param(":name")
	busService := new(service.BusService)
	bus := busService.GetOne(name)

	busData, error := json.Marshal(bus)
	if error != nil {
		println(error)
	}

	result := string(busData)
	println(result)

	this.Ctx.WriteString(result)
}

func (this BusController) Add() {
	station := this.Ctx.Input.Param(":station")
	busService := new(service.BusService)
	bus := models.Bus{Station: station, Num: 10, Count: 1, Status: "Enable"}
	busService.GetOn(bus)

	buses := busService.GetOne(station)
	r := new(models.R)
	r.Msg = " add successfully"
	r.Code = 200
	r.Data = buses
	rData, error := json.Marshal(r)
	if error != nil {
		println(error)
	}

	result := string(rData)
	println(result)
	this.Ctx.WriteString(result)
}

func (this BusController) Delete() {
	id := this.Ctx.Input.Param(":id")
	println(id)

	busService := new(service.BusService)
	busService.Stop(id)

	r := new(models.R)
	r.Msg = id + " delete successfully"
	r.Code = 200

	rData, error := json.Marshal(r)
	if error != nil {
		println(error)
	}

	result := string(rData)

	println(result)

	this.Ctx.WriteString(result)
}
