package apis

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/go-admin-team/go-admin-core/sdk/api"
	"github.com/go-admin-team/go-admin-core/sdk/pkg/jwtauth/user"
	_ "github.com/go-admin-team/go-admin-core/sdk/pkg/response"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
)

type Applications struct {
	api.Api
}

// GetPage 获取Applications列表
// @Summary 获取Applications列表
// @Description 获取Applications列表
// @Tags Applications
// @Param title query string false "标题"
// @Param status query int64 false "状态"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.Applications}} "{"code": 200, "data": [...]}"
// @Router /api/v1/applications [get]
// @Security Bearer
func (e Applications) GetPage(c *gin.Context) {
	req := dto.ApplicationsGetPageReq{}
	s := service.Applications{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	p := actions.GetPermissionFromContext(c)
	list := make([]models.Applications, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Applications失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取Applications
// @Summary 获取Applications
// @Description 获取Applications
// @Tags Applications
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.Applications} "{"code": 200, "data": [...]}"
// @Router /api/v1/applications/{id} [get]
// @Security Bearer
func (e Applications) Get(c *gin.Context) {
	req := dto.ApplicationsGetReq{}
	s := service.Applications{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	var object models.Applications

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取Applications失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建Applications
// @Summary 创建Applications
// @Description 创建Applications
// @Tags Applications
// @Accept application/json
// @Product application/json
// @Param data body dto.ApplicationsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/applications [post]
// @Security Bearer
func (e Applications) Insert(c *gin.Context) {
	req := dto.ApplicationsInsertReq{}
	s := service.Applications{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	// 设置创建人
	req.SetCreateBy(user.GetUserId(c))

	err = s.Insert(&req)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("创建Applications失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改Applications
// @Summary 修改Applications
// @Description 修改Applications
// @Tags Applications
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ApplicationsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/applications/{id} [put]
// @Security Bearer
func (e Applications) Update(c *gin.Context) {
	req := dto.ApplicationsUpdateReq{}
	s := service.Applications{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}
	req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Update(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("修改Applications失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除Applications
// @Summary 删除Applications
// @Description 删除Applications
// @Tags Applications
// @Param data body dto.ApplicationsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/applications [delete]
// @Security Bearer
func (e Applications) Delete(c *gin.Context) {
	s := service.Applications{}
	req := dto.ApplicationsDeleteReq{}
	err := e.MakeContext(c).
		MakeOrm().
		Bind(&req).
		MakeService(&s.Service).
		Errors
	if err != nil {
		e.Logger.Error(err)
		e.Error(500, err, err.Error())
		return
	}

	// req.SetUpdateBy(user.GetUserId(c))
	p := actions.GetPermissionFromContext(c)

	err = s.Remove(&req, p)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("删除Applications失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
