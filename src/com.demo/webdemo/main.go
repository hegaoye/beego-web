package main

import (
	"com.demo/webdemo/controller"
	"com.demo/webdemo/rpcserver"
	beego "github.com/beego/beego/v2/adapter"
)

func main() {
	go rpcserver.StartRpcServer(18003)

	//job.Job()

	beego.ErrorController(&controller.ErrorController{})

	apiRouter := beego.NewNamespace("/api",
		beego.NSRouter("/get/:name", &controller.BusController{}, "get:Get"),
		beego.NSRouter("/add/:station", &controller.BusController{}, "get:Add"),
		beego.NSRouter("/delete/:id", &controller.BusController{}, "get:Delete"),
		beego.NSRouter("/app/all", &controller.AppController{}, "get:ListAll"))

	userRouter := beego.NewNamespace("/user",
		beego.NSRouter("/all", &controller.UserController{}, "get:ListAll"),
		beego.NSRouter("/get/name/:name", &controller.UserController{}, "get:GetOneByName"),
		beego.NSRouter("/get/workno/:workNo", &controller.UserController{}, "get:GetOneByWorkNo"),
		beego.NSRouter("/update/name/:name/:u", &controller.UserController{}, "get:UpdateUByName"),
		beego.NSRouter("/upload", &controller.UserController{}, "post:UploadData"))

	redisRouter := beego.NewNamespace("/redis",
		beego.NSRouter("/get/:key", &controller.RedisController{}, "get:Get"),
		beego.NSRouter("/add/:key/:value", &controller.RedisController{}, "get:Add"),
		beego.NSRouter("/del/:key", &controller.RedisController{}, "get:Del"))

	beego.Router("/heartbeat", &controller.HeartbeatController{}, "get:Heartbeat")
	beego.Router("/rpc/:data", &controller.GrpcTestController{}, "get:SendMsg")

	beego.AddNamespace(apiRouter, userRouter, redisRouter)

	beego.BConfig.CopyRequestBody = true
	beego.BConfig.RunMode = beego.DEV
	beego.BConfig.Listen.HTTPPort = 8080

	beego.Run()
}
