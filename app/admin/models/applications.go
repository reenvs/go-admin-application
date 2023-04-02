package models

import (
	"go-admin/common/models"
)

type Applications struct {
	models.Model

	Title       string `json:"title" gorm:"type:varchar(128);comment:标题"`
	Description string `json:"description" gorm:"type:varchar(255);comment:描述"`
	Status      int64  `json:"status" gorm:"type:int(1);comment:状态"`
	models.ModelTime
	models.ControlBy
}

func (Applications) TableName() string {
	return "applications"
}

func (e *Applications) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *Applications) GetId() interface{} {
	return e.Id
}
