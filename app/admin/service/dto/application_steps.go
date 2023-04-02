package dto

import (

	//"time"

	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ApplicationStepsGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:contains;column:title;table:application_steps" comment:"标题"`
	Type           int64  `form:"type"  search:"type:exact;column:type;table:application_steps" comment:"类型，1:部门，2:个人"`
	Step           int64  `form:"step"  search:"type:exact;column:step;table:application_steps" comment:"审批步骤"`
	Level          int64  `form:"level"  search:"type:exact;column:level;table:application_steps" comment:"部门级别，0代表级，1代表上一级，2代表上两级"`
	SysUserId      int64  `form:"sysUserId"  search:"type:exact;column:sys_user_id;table:application_steps" comment:"用户ID"`
	ApplicationId  int64  `form:"applicationId"  search:"type:exact;column:application_id;table:application_steps" comment:"关联审批ID"`
	ApplicationStepsOrder
}

type ApplicationStepsOrder struct {
	Id            string `form:"idOrder"  search:"type:order;column:id;table:application_steps"`
	Title         string `form:"titleOrder"  search:"type:order;column:title;table:application_steps"`
	Description   string `form:"descriptionOrder"  search:"type:order;column:description;table:application_steps"`
	Type          string `form:"typeOrder"  search:"type:order;column:type;table:application_steps"`
	Step          string `form:"stepOrder"  search:"type:order;column:step;table:application_steps"`
	Level         string `form:"levelOrder"  search:"type:order;column:level;table:application_steps"`
	SysUserId     string `form:"sysUserIdOrder"  search:"type:order;column:sys_user_id;table:application_steps"`
	CreatedAt     string `form:"createdAtOrder"  search:"type:order;column:created_at;table:application_steps"`
	UpdatedAt     string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:application_steps"`
	DeletedAt     string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:application_steps"`
	CreateBy      string `form:"createByOrder"  search:"type:order;column:create_by;table:application_steps"`
	UpdateBy      string `form:"updateByOrder"  search:"type:order;column:update_by;table:application_steps"`
	ApplicationId string `form:"applicationIdOrder"  search:"type:order;column:application_id;table:application_steps"`
}

func (m *ApplicationStepsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ApplicationStepsInsertReq struct {
	Id            int    `json:"-" comment:"主键编码"` // 主键编码
	Title         string `json:"title" comment:"标题"`
	Description   string `json:"description" comment:"描述"`
	Type          int64  `json:"type" comment:"类型，1:部门，2:个人"`
	Step          int64  `json:"step" comment:"审批步骤"`
	Level         int64  `json:"level" comment:"部门级别，0代表级，1代表上一级，2代表上两级"`
	SysUserId     int64  `json:"sysUserId" comment:"用户ID"`
	ApplicationId int64  `json:"applicationId" comment:"关联审批ID"`
	common.ControlBy
}

func (s *ApplicationStepsInsertReq) Generate(model *models.ApplicationSteps) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Description = s.Description
	model.Type = s.Type
	model.Step = s.Step
	model.Level = s.Level
	model.SysUserId = s.SysUserId
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
	model.ApplicationId = s.ApplicationId
}

func (s *ApplicationStepsInsertReq) GetId() interface{} {
	return s.Id
}

type ApplicationStepsUpdateReq struct {
	Id            int    `uri:"id" comment:"主键编码"` // 主键编码
	Title         string `json:"title" comment:"标题"`
	Description   string `json:"description" comment:"描述"`
	Type          int64  `json:"type" comment:"类型，1:部门，2:个人"`
	Step          int64  `json:"step" comment:"审批步骤"`
	Level         int64  `json:"level" comment:"部门级别，0代表级，1代表上一级，2代表上两级"`
	SysUserId     int64  `json:"sysUserId" comment:"用户ID"`
	ApplicationId int64  `json:"applicationId" comment:"关联审批ID"`
	common.ControlBy
}

func (s *ApplicationStepsUpdateReq) Generate(model *models.ApplicationSteps) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Description = s.Description
	model.Type = s.Type
	model.Step = s.Step
	model.Level = s.Level
	model.SysUserId = s.SysUserId
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
	model.ApplicationId = s.ApplicationId
}

func (s *ApplicationStepsUpdateReq) GetId() interface{} {
	return s.Id
}

// ApplicationStepsGetReq 功能获取请求参数
type ApplicationStepsGetReq struct {
	Id int `uri:"id"`
}

func (s *ApplicationStepsGetReq) GetId() interface{} {
	return s.Id
}

// ApplicationStepsDeleteReq 功能删除请求参数
type ApplicationStepsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ApplicationStepsDeleteReq) GetId() interface{} {
	return s.Ids
}
