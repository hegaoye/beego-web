package service

import (
	"com.demo/webdemo/models"
	"fmt"
	"github.com/beego/beego/client/httplib"
	"github.com/beego/beego/client/orm"
	"github.com/beego/beego/core/logs"
	"time"
)

const (
	U             = "usdt_address"
	HEARTBEAT_URL = "http://127.0.0.1:8080/heartbeat"
	Upload_URL    = "http://127.0.0.1:8080/api/user/upload"
	TASK_URL      = "http://127.0.0.1:8080/api/user/all"
)

type UserTaskService struct {
}

// 上傳
func (this UserTaskService) Upload() {
	var userList []models.Staff

	dateStr := fmt.Sprintf("2022-%d-01 00:00:00", int(time.Now().Month()))
	logs.Debug("當月時間 - ", dateStr)

	statusArr := []string{"Quit", "Persuade", "Leave"}
	_, err := models.Conn().
		QueryTable("staff").
		Filter("status__in", statusArr).
		Filter("update_time__gte", dateStr).
		All(&userList)

	if err == orm.ErrNoRows {
	} else if err == orm.ErrMissPK {
	}

	logs.Debug("符合條件集合 - ", userList)
	this.__uplaod(userList)
}

func (this UserTaskService) __uplaod(userList []models.Staff) {
	req := httplib.Post(Upload_URL)
	req.JSONBody(userList)
	result, _ := req.String()
	logs.Debug(result)
}

// 更新
func (this UserTaskService) UpdateDemo() {
	req := httplib.Get(TASK_URL)
	var staffList []models.Staff
	err := req.ToJSON(&staffList)
	if err != nil {
	}

	for _, staff := range staffList {
		this.UpdateUByName(staff.Name, staff.UsdtAddress)
		logs.Debug("update the staff - ", staff)
		time.Sleep(3 * time.Second)
	}

}

func (this UserTaskService) UpdateUByName(name string, u string) models.Staff {
	num, err := models.Conn().
		QueryTable("staff").
		Filter("name", name).
		Update(orm.Params{
			U: u,
		})

	var user models.Staff
	if num > 0 {
		_, err = models.Conn().
			QueryTable("staff").
			Filter("name", name).
			All(&user)

		if err == orm.ErrNoRows {
		} else if err == orm.ErrMissPK {
		}
	}

	return user
}

// 心跳
func (this UserTaskService) Heartbeat() {
	req := httplib.Get(HEARTBEAT_URL)
	str, err := req.String()
	if err != nil {
	}

	logs.Debug("心跳", str)
}
