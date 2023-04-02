package dto

import (
	"go-admin/app/admin/models"
	"go-admin/common/dto"
	common "go-admin/common/models"
)

type ApplicationsGetPageReq struct {
	dto.Pagination `search:"-"`
	Title          string `form:"title"  search:"type:contains;column:title;table:applications" comment:"标题"`
	Status         int64  `form:"status"  search:"type:exact;column:status;table:applications" comment:"状态"`
	ApplicationsOrder
}

type ApplicationsOrder struct {
	Id          string `form:"idOrder"  search:"type:order;column:id;table:applications"`
	Title       string `form:"titleOrder"  search:"type:order;column:title;table:applications"`
	Description string `form:"descriptionOrder"  search:"type:order;column:description;table:applications"`
	Status      string `form:"statusOrder"  search:"type:order;column:status;table:applications"`
	CreatedAt   string `form:"createdAtOrder"  search:"type:order;column:created_at;table:applications"`
	UpdatedAt   string `form:"updatedAtOrder"  search:"type:order;column:updated_at;table:applications"`
	DeletedAt   string `form:"deletedAtOrder"  search:"type:order;column:deleted_at;table:applications"`
	CreateBy    string `form:"createByOrder"  search:"type:order;column:create_by;table:applications"`
	UpdateBy    string `form:"updateByOrder"  search:"type:order;column:update_by;table:applications"`
}

func (m *ApplicationsGetPageReq) GetNeedSearch() interface{} {
	return *m
}

type ApplicationsInsertReq struct {
	Id          int    `json:"-" comment:"主键编码"` // 主键编码
	Title       string `json:"title" comment:"标题"`
	Description string `json:"description" comment:"描述"`
	Status      int64  `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ApplicationsInsertReq) Generate(model *models.Applications) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Description = s.Description
	model.Status = s.Status
	model.CreateBy = s.CreateBy // 添加这而，需要记录是被谁创建的
}

func (s *ApplicationsInsertReq) GetId() interface{} {
	return s.Id
}

type ApplicationsUpdateReq struct {
	Id          int    `uri:"id" comment:"主键编码"` // 主键编码
	Title       string `json:"title" comment:"标题"`
	Description string `json:"description" comment:"描述"`
	Status      int64  `json:"status" comment:"状态"`
	common.ControlBy
}

func (s *ApplicationsUpdateReq) Generate(model *models.Applications) {
	if s.Id == 0 {
		model.Model = common.Model{Id: s.Id}
	}
	model.Title = s.Title
	model.Description = s.Description
	model.Status = s.Status
	model.UpdateBy = s.UpdateBy // 添加这而，需要记录是被谁更新的
}

func (s *ApplicationsUpdateReq) GetId() interface{} {
	return s.Id
}

// ApplicationsGetReq 功能获取请求参数
type ApplicationsGetReq struct {
	Id int `uri:"id"`
}

func (s *ApplicationsGetReq) GetId() interface{} {
	return s.Id
}

// ApplicationsDeleteReq 功能删除请求参数
type ApplicationsDeleteReq struct {
	Ids []int `json:"ids"`
}

func (s *ApplicationsDeleteReq) GetId() interface{} {
	return s.Ids
}
