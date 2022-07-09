package models

import "github.com/beego/beego/client/orm"

type App struct {
	Id         string
	Name       string
	DriverCode string
	Summary    string
}

func init() {
	orm.RegisterModel(new(App))
}
