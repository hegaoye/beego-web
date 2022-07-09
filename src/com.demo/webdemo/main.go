package main

import (
	"com.demo/webdemo/controller"
	beego "github.com/beego/beego/adapter"
)

func main() {
	//job.Job()
	beego.ErrorController(&controller.ErrorController{})
	beego.Router("/api/get/:name", &controller.BusController{}, "get:Get")
	beego.Router("/api/add/:station", &controller.BusController{}, "get:Add")
	beego.Router("/api/delete/:id", &controller.BusController{}, "get:Delete")
	beego.Router("/api/app/all", &controller.AppController{}, "get:ListAll")
	beego.Router("/api/redis/get/:key", &controller.RedisController{}, "get:Get")
	beego.Router("/api/redis/add/:key/:value", &controller.RedisController{}, "get:Add")
	beego.Router("/api/redis/del/:key", &controller.RedisController{}, "get:Del")
	beego.Run()
}
