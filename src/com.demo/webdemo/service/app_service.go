package service

import (
	"com.demo/webdemo/models"
	"fmt"
	"github.com/beego/beego/v2/client/orm"
)

type AppService struct {
}

func (this AppService) ListAll() []models.App {
	var appList []models.App

	_, err := models.Conn().
		QueryTable("app").
		All(&appList)

	if err == orm.ErrNoRows {
		fmt.Println("查询不到")
	} else if err == orm.ErrMissPK {
		fmt.Println("找不到主键")
	}

	return appList
}
