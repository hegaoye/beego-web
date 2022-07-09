package models

import "gorm.io/gorm"

// 公交車
type Bus struct {
	gorm.Model
	Num     int8   `json:"num"`
	Station string `json:"station"`
	Count   int    `json:"count"`
	Status  string `json:"status"`
}
