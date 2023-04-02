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

type Applications struct {
	service.Service
}

// GetPage 获取Applications列表
func (e *Applications) GetPage(c *dto.ApplicationsGetPageReq, p *actions.DataPermission, list *[]models.Applications, count *int64) error {
	var err error
	var data models.Applications

	err = e.Orm.Model(&data).
		Scopes(
			cDto.MakeCondition(c.GetNeedSearch()),
			cDto.Paginate(c.GetPageSize(), c.GetPageIndex()),
			actions.Permission(data.TableName(), p),
		).
		Find(list).Limit(-1).Offset(-1).
		Count(count).Error
	if err != nil {
		e.Log.Errorf("ApplicationsService GetPage error:%s \r\n", err)
		return err
	}
	return nil
}

// Get 获取Applications对象
func (e *Applications) Get(d *dto.ApplicationsGetReq, p *actions.DataPermission, model *models.Applications) error {
	var data models.Applications

	err := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).
		First(model, d.GetId()).Error
	if err != nil && errors.Is(err, gorm.ErrRecordNotFound) {
		err = errors.New("查看对象不存在或无权查看")
		e.Log.Errorf("Service GetApplications error:%s \r\n", err)
		return err
	}
	if err != nil {
		e.Log.Errorf("db error:%s", err)
		return err
	}
	return nil
}

// Insert 创建Applications对象
func (e *Applications) Insert(c *dto.ApplicationsInsertReq) error {
	var err error
	var data models.Applications
	c.Generate(&data)
	err = e.Orm.Create(&data).Error
	if err != nil {
		e.Log.Errorf("ApplicationsService Insert error:%s \r\n", err)
		return err
	}

	stepX := &ApplicationSteps{service.Service{Orm: e.Orm}}
	err = stepX.Insert(&dto.ApplicationStepsInsertReq{
		Id:            0,
		Title:         data.Title,
		Description:   data.Description,
		Type:          0,
		Step:          0,
		Level:         0,
		SysUserId:     int64(data.CreateBy),
		ApplicationId: int64(data.Id),
		ControlBy:     data.ControlBy,
	})
	if err != nil {
		e.Log.Errorf("ApplicationsService step1 Insert error:%s \r\n", err)
		return err
	}

	err = stepX.Insert(&dto.ApplicationStepsInsertReq{
		Id:            0,
		Title:         data.Title,
		Description:   data.Description,
		Type:          1,
		Step:          1,
		Level:         1,
		SysUserId:     int64(data.CreateBy),
		ApplicationId: int64(data.Id),
		ControlBy:     data.ControlBy,
	})
	if err != nil {
		e.Log.Errorf("ApplicationsService step1 Insert error:%s \r\n", err)
		return err
	}

	return nil
}

// Update 修改Applications对象
func (e *Applications) Update(c *dto.ApplicationsUpdateReq, p *actions.DataPermission) error {
	var err error
	var data = models.Applications{}
	e.Orm.Scopes(
		actions.Permission(data.TableName(), p),
	).First(&data, c.GetId())
	c.Generate(&data)

	db := e.Orm.Save(&data)
	if err = db.Error; err != nil {
		e.Log.Errorf("ApplicationsService Save error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权更新该数据")
	}
	return nil
}

// Remove 删除Applications
func (e *Applications) Remove(d *dto.ApplicationsDeleteReq, p *actions.DataPermission) error {
	var data models.Applications

	db := e.Orm.Model(&data).
		Scopes(
			actions.Permission(data.TableName(), p),
		).Delete(&data, d.GetId())
	if err := db.Error; err != nil {
		e.Log.Errorf("Service RemoveApplications error:%s \r\n", err)
		return err
	}
	if db.RowsAffected == 0 {
		return errors.New("无权删除该数据")
	}
	return nil
}
