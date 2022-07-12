package models

import (
	"github.com/beego/beego/v2/adapter/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	MYSQL_URL = "{username}:{password}@tcp({host}:{port})/{database}?charset=utf8"
)

var session orm.Ormer

func Conn() orm.Ormer {
	return session
}

/*
*
database = payment_dev
username = payment_dev
password = TEkAgPyL1GkaUk0h
host = zt.cg45.xyz
port = 31008
*/
func init() {
	//数据库名称
	database := "payment_dev"
	//database := beego.AppConfig.String("database")
	//数据库连接用户名
	username := "payment_dev"
	//username := beego.AppConfig.String("username")
	//数据库连接用户名
	password := "TEkAgPyL1GkaUk0h"
	//数据库IP（域名）
	host := "zt.cg45.xyz"
	//host := "basic-tidb.public"
	//数据库端口
	port := "31008"
	//port := "4000"

	url := MYSQL_URL
	url = strings.Replace(url, "{username}", username, -1)
	url = strings.Replace(url, "{password}", password, -1)
	url = strings.Replace(url, "{host}", host, -1)
	url = strings.Replace(url, "{port}", port, -1)
	url = strings.Replace(url, "{database}", database, -1)

	//連接數據庫
	orm.RegisterDriver("mysql", orm.DRMySQL)
	orm.RegisterDataBase("default", "mysql", url)

	//打開連接
	session = orm.NewOrm()
}
