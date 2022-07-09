package service

import "com.demo/webdemo/models"

type BusService struct {
}

func (this BusService) GetOne(station string) []models.Bus {
	var buses []models.Bus
	Conn().Where("station = ?", station).Find(&buses)
	return buses
}

func (this BusService) GetOn(bus models.Bus) bool {
	Conn().Create(&bus)
	return true
}

func (this BusService) GetOff(bus models.Bus) bool {
	Conn().Delete(&bus)
	return true
}

func (this BusService) Count(busNum int) int64 {
	var count int64
	Conn().Model(&models.Bus{}).
		Where("num>?", busNum).
		Count(&count)
	return count
}

func (this BusService) Stop(id string) bool {
	Conn().Model(&models.Bus{}).
		Where("id = ?", id).
		Update("status", "Disable")
	return true
}
