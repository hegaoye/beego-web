package service

import (
	"com.demo/webdemo/models"
	"gorm.io/gorm"
)

type BusDao interface {
	GetOne(station string) []models.Bus
	GetOn(bus models.Bus) bool
	GetOff(bus models.Bus) bool
	Count(busNum int) int64
	Stop(db *gorm.DB) bool
}
