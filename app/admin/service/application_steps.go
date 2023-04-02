package service

import (
	"errors"

	"github.com/go-admin-team/go-admin-core/sdk/service"
	"gorm.io/gorm"

	"go-admin/app/admin/models"
	"go-admin/app/admin/service/dto"
	"go-admin/common/actions"
	cDto "go-admin/common/dto"
)

type ApplicationSteps struct {
	service.Service
}

// GetPage 获取ApplicationSteps列表
func (e *ApplicationSteps) GetPage(c *dto.ApplicationStepsGetPageReq, p *actions.DataPermission, list *[]models.ApplicationSteps, count *int64) error {
	var err error
	var data models.ApplicationSteps

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ApplicationStepsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取ApplicationSteps对象
func (e *ApplicationSteps) Get(d *dto.ApplicationStepsGetReq, p *actions.DataPermission, model *models.ApplicationSteps) error {
	var data models.ApplicationSteps

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetApplicationSteps error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建ApplicationSteps对象
func (e *ApplicationSteps) Insert(c *dto.ApplicationStepsInsertReq) error {
	var err error
	var data models.ApplicationSteps
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ApplicationStepsService Insert error:%s \r\n", err)
		return err
	}
	return nil
}

// Update 修改ApplicationSteps对象
func (e *ApplicationSteps) Update(c *dto.ApplicationStepsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.ApplicationSteps{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ApplicationStepsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除ApplicationSteps
func (e *ApplicationSteps) Remove(d *dto.ApplicationStepsDeleteReq, p *actions.DataPermission) error {
	var data models.ApplicationSteps

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveApplicationSteps error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
