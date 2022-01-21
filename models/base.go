package models

import (
	"github.com/GoAdminGroup/go-admin/modules/db"
	"github.com/jinzhu/gorm"
)

var (
	Orm *gorm.DB
	err error
)

func Init(c db.Connection) {
	Orm, err = gorm.Open("postgres", c.GetDB("default"))

	if err != nil {
		panic("initialize orm failed")
	}
	print(Orm)
}
