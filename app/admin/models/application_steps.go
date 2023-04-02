package models

import (

	//"time"

	"go-admin/common/models"
)

type ApplicationSteps struct {
	models.Model

	Title         string `json:"title" gorm:"type:varchar(128);comment:标题"`
	Description   string `json:"description" gorm:"type:varchar(128);comment:描述"`
	Type          int64  `json:"type" gorm:"type:int(11);comment:类型，1:部门，2:个人"`
	Step          int64  `json:"step" gorm:"type:int(11);comment:审批步骤"`
	Level         int64  `json:"level" gorm:"type:int(11);comment:部门级别，0代表级，1代表上一级，2代表上两级"`
	SysUserId     int64  `json:"sysUserId" gorm:"type:int(11);comment:用户ID"`
	ApplicationId int64  `json:"applicationId" gorm:"type:int(11);comment:关联审批ID"`
	models.ModelTime
	models.ControlBy
}

func (ApplicationSteps) TableName() string {
	return "application_steps"
}

func (e *ApplicationSteps) Generate() models.ActiveRecord {
	o := *e
	return &o
}

func (e *ApplicationSteps) GetId() interface{} {
	return e.Id
}
