package main

import (
	"com.demo/webdemo/controller"
	"com.demo/webdemo/job"
	beego "github.com/beego/beego/adapter"
)

func main() {
	job.Job()
	beego.ErrorController(&controller.ErrorController{})
	beego.Router("/api/get/:name", &controller.BusController{}, "get:Get")
	beego.Router("/api/add/:station", &controller.BusController{}, "get:Add")
	beego.Router("/api/delete/:id", &controller.BusController{}, "get:Delete")
	beego.Router("/api/app/all", &controller.AppController{}, "get:ListAll")

	beego.Router("/api/user/all", &controller.UserController{}, "get:ListAll")
	beego.Router("/api/user/get/name/:name", &controller.UserController{}, "get:GetOneByName")
	beego.Router("/api/user/get/workno/:workNo", &controller.UserController{}, "get:GetOneByWorkNo")
	beego.Router("/api/user/update/name/:name/:u", &controller.UserController{}, "get:UpdateUByName")
	beego.Router("/api/user/upload", &controller.UserController{}, "post:UploadData")

	beego.Router("/heartbeat", &controller.HeartbeatController{}, "get:Heartbeat")

	beego.Router("/api/redis/get/:key", &controller.RedisController{}, "get:Get")
	beego.Router("/api/redis/add/:key/:value", &controller.RedisController{}, "get:Add")
	beego.Router("/api/redis/del/:key", &controller.RedisController{}, "get:Del")

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.RunMode = beego.DEV

	beego.Run()
}
