package service

import (
	"com.demo/webdemo/models"
	"github.com/beego/beego/client/orm"
)

type UserService struct {
}

func (this UserService) ListAll() []models.Staff {
	var appList []models.Staff

	_, err := models.Conn().
		QueryTable("staff").
		All(&appList)

	if err == orm.ErrNoRows {

	} else if err == orm.ErrMissPK {

	}

	return appList
}

func (this UserService) GetOneByName(name string) models.Staff {
	var user models.Staff
	_, err := models.Conn().
		QueryTable("staff").
		Filter("name", name).
		All(&user)

	if err == orm.ErrNoRows {

	} else if err == orm.ErrMissPK {

	}

	return user
}

func (this UserService) GetOneByWorkNo(workNo string) models.Staff {
	var user models.Staff
	_, err := models.Conn().
		QueryTable("staff").
		Filter("work_no", workNo).
		All(&user)

	if err == orm.ErrNoRows {

	} else if err == orm.ErrMissPK {

	}

	return user
}

func (this UserService) UpdateUByName(name string, u string) models.Staff {
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
