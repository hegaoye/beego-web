package main

import (
	"com.demo/webdemo/controller"
	"com.demo/webdemo/job"
	beego "github.com/beego/beego/adapter"
)

func main() {
	job.Job()

	beego.ErrorController(&controller.ErrorController{})

	userRouter := beego.NewNamespace("/user",
		beego.NSRouter("/all", &controller.UserController{}, "get:ListAll"),
		beego.NSRouter("/get/name/:name", &controller.UserController{}, "get:GetOneByName"),
		beego.NSRouter("/get/workno/:workNo", &controller.UserController{}, "get:GetOneByWorkNo"),
		beego.NSRouter("/update/name/:name/:u", &controller.UserController{}, "get:UpdateUByName"),
		beego.NSRouter("/upload", &controller.UserController{}, "post:UploadData"))

	beego.Router("/heartbeat", &controller.HeartbeatController{}, "get:Heartbeat")

	beego.AddNamespace(userRouter)
	beego.BConfig.CopyRequestBody = true
	beego.BConfig.RunMode = beego.DEV

	beego.Run()
}
