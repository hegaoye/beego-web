package models

import (
	beego "github.com/beego/beego/v2/adapter"
	"github.com/beego/beego/v2/client/orm"
	_ "github.com/go-sql-driver/mysql"
	"strings"
)

const (
	MysqlUrl = "{username}:{password}@tcp({host}:{port})/{database}?charset=utf8"
)

var session orm.Ormer

func Conn() orm.Ormer {
	return session
}

func init() {
	//数据库名称
	database := beego.AppConfig.String("database")
	//数据库连接用户名
	username := beego.AppConfig.String("username")
	//数据库连接用户名
	password := beego.AppConfig.String("password")
	//数据库IP（域名）
	host := beego.AppConfig.String("host")
	//数据库端口
	port := beego.AppConfig.String("port")

	url := MysqlUrl
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
