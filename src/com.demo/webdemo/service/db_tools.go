package service

import "gorm.io/gorm"
import "gorm.io/driver/sqlite"

var session *gorm.DB

// 获取数据库连接
func Conn() *gorm.DB {
	return session
}

func init() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		println("error......")
	}
	session = db
}
