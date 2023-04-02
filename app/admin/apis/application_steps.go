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

type ApplicationSteps struct {
	api.Api
}

// GetPage 获取ApplicationSteps列表
// @Summary 获取ApplicationSteps列表
// @Description 获取ApplicationSteps列表
// @Tags ApplicationSteps
// @Param title query string false "标题"
// @Param type query int64 false "类型，1:部门，2:个人"
// @Param step query int64 false "审批步骤"
// @Param level query int64 false "部门级别，0代表级，1代表上一级，2代表上两级"
// @Param sysUserId query int64 false "用户ID"
// @Param applicationId query int64 false "关联审批ID"
// @Param pageSize query int false "页条数"
// @Param pageIndex query int false "页码"
// @Success 200 {object} response.Response{data=response.Page{list=[]models.ApplicationSteps}} "{"code": 200, "data": [...]}"
// @Router /api/v1/application-steps [get]
// @Security Bearer
func (e ApplicationSteps) GetPage(c *gin.Context) {
	req := dto.ApplicationStepsGetPageReq{}
	s := service.ApplicationSteps{}
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
	list := make([]models.ApplicationSteps, 0)
	var count int64

	err = s.GetPage(&req, p, &list, &count)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ApplicationSteps失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.PageOK(list, int(count), req.GetPageIndex(), req.GetPageSize(), "查询成功")
}

// Get 获取ApplicationSteps
// @Summary 获取ApplicationSteps
// @Description 获取ApplicationSteps
// @Tags ApplicationSteps
// @Param id path int false "id"
// @Success 200 {object} response.Response{data=models.ApplicationSteps} "{"code": 200, "data": [...]}"
// @Router /api/v1/application-steps/{id} [get]
// @Security Bearer
func (e ApplicationSteps) Get(c *gin.Context) {
	req := dto.ApplicationStepsGetReq{}
	s := service.ApplicationSteps{}
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
	var object models.ApplicationSteps

	p := actions.GetPermissionFromContext(c)
	err = s.Get(&req, p, &object)
	if err != nil {
		e.Error(500, err, fmt.Sprintf("获取ApplicationSteps失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(object, "查询成功")
}

// Insert 创建ApplicationSteps
// @Summary 创建ApplicationSteps
// @Description 创建ApplicationSteps
// @Tags ApplicationSteps
// @Accept application/json
// @Product application/json
// @Param data body dto.ApplicationStepsInsertReq true "data"
// @Success 200 {object} response.Response	"{"code": 200, "message": "添加成功"}"
// @Router /api/v1/application-steps [post]
// @Security Bearer
func (e ApplicationSteps) Insert(c *gin.Context) {
	req := dto.ApplicationStepsInsertReq{}
	s := service.ApplicationSteps{}
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
		e.Error(500, err, fmt.Sprintf("创建ApplicationSteps失败，\r\n失败信息 %s", err.Error()))
		return
	}

	e.OK(req.GetId(), "创建成功")
}

// Update 修改ApplicationSteps
// @Summary 修改ApplicationSteps
// @Description 修改ApplicationSteps
// @Tags ApplicationSteps
// @Accept application/json
// @Product application/json
// @Param id path int true "id"
// @Param data body dto.ApplicationStepsUpdateReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "修改成功"}"
// @Router /api/v1/application-steps/{id} [put]
// @Security Bearer
func (e ApplicationSteps) Update(c *gin.Context) {
	req := dto.ApplicationStepsUpdateReq{}
	s := service.ApplicationSteps{}
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
		e.Error(500, err, fmt.Sprintf("修改ApplicationSteps失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "修改成功")
}

// Delete 删除ApplicationSteps
// @Summary 删除ApplicationSteps
// @Description 删除ApplicationSteps
// @Tags ApplicationSteps
// @Param data body dto.ApplicationStepsDeleteReq true "body"
// @Success 200 {object} response.Response	"{"code": 200, "message": "删除成功"}"
// @Router /api/v1/application-steps [delete]
// @Security Bearer
func (e ApplicationSteps) Delete(c *gin.Context) {
	s := service.ApplicationSteps{}
	req := dto.ApplicationStepsDeleteReq{}
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
		e.Error(500, err, fmt.Sprintf("删除ApplicationSteps失败，\r\n失败信息 %s", err.Error()))
		return
	}
	e.OK(req.GetId(), "删除成功")
}
