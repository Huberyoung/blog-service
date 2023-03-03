package v1

import (
	"blog-service/global"
	"blog-service/interbal/model"
	"blog-service/interbal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/convert"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Tag struct {
}

func NewTage() Tag {
	return Tag{}
}

// Get godoc
//
//	@Summary		获取单个标签
//	@Description	通过id获取单个标签
//	@Tags			Tag
//	@Produce		json
//	@Param			id	path		int				true	"标签ID"
//	@Success		200	{object}	model.Tag		"成功"
//	@Failure		400	{object}	errcode.Error	"请求错误"
//	@Failure		404	{object}	errcode.Error	"找不到页面"
//	@Failure		500	{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/tags/:id [get]
func (t Tag) Get(c *gin.Context) {
	var tag model.Tag
	global.DBEngine.First(&tag)
	log.Printf("record %#v\n", tag)

	c.JSON(http.StatusOK, tag)
}

// List godoc
//
//	@Summary		获取多个标签
//	@Description	通过获取多个标签
//	@Tags			Tag
//	@Produce		json
//	@Param			name		query		string			false	"标签名称"	maxlength(100)
//	@Param			state		query		int				false	"状态"	Enums(0, 1)	default(1)
//	@Param			page		query		int				false	"页码"
//	@Param			page_size	query		int				false	"每页数量"
//	@Success		200			{object}	model.Tag		"成功"
//	@Failure		400			{object}	errcode.Error	"请求错误"
//	@Failure		404			{object}	errcode.Error	"找不到页面"
//	@Failure		500			{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}

	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errors)
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.ErrorF("svc.CountTag err:%v", err)
		response.ToErrorResponseList(errcode.ErrorCountTagListFail)
		return
	}

	tags, err := svc.ListTag(&param, &pager)
	if err != nil {
		global.Logger.ErrorF("svc.ListTag err:%v", err)
		response.ToErrorResponseList(errcode.ErrorGetTagListFail)
		return
	}
	response.ToResponseList(tags, totalRows)
	return
}

// Create godoc
//
//	@Summary		新增标签
//	@Description	新增标签
//	@Tags			Tag
//	@Produce		json
//	@Param			name		body		string			true	"标签名称"	minlength(3)	maxlength(100)
//	@Param			state		body		int				false	"状态"	Enums(0, 1)		default(1)
//	@Param			created_by	body		string			true	"创建者"	minlength(3)	maxlength(100)
//	@Success		200			{object}	model.Tag		"成功"
//	@Failure		400			{object}	errcode.Error	"请求错误"
//	@Failure		404			{object}	errcode.Error	"找不到页面"
//	@Failure		500			{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}

	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%s", errors.Error())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.ErrorF("app.CreateTag errs:%v", errors)
		response.ToErrorResponseList(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Update godoc
//
//	@Summary		获取多个标签
//	@Description	通过id获取多个标签
//	@Tags			Tag
//	@Produce		json
//	@Param			id			path		int				true	"标签ID"
//	@Param			name		body		string			true	"标签名称"	minlength(3)	maxlength(100)
//	@Param			state		body		int				false	"状态"	Enums(0, 1)		default(1)
//	@Param			modified_by	body		string			true	"修改者"	minlength(3)	maxlength(100)
//	@Success		200			{object}	model.Tag		"成功"
//	@Failure		400			{object}	errcode.Error	"请求错误"
//	@Failure		404			{object}	errcode.Error	"找不到页面"
//	@Failure		500			{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/tags/:id [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt()}

	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%s", errors.Errors())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.ErrorF("app.UpdateTag errs:%v", errors)
		response.ToErrorResponseList(errcode.ErrorUpdateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Delete godoc
//
//	@Summary		删除标签
//	@Description	通过id删除标签
//	@Tags			Tag
//	@Produce		json
//	@Param			id	path		int				true	"标签ID"
//	@Success		200	{object}	model.Tag		"成功"
//	@Failure		400	{object}	errcode.Error	"请求错误"
//	@Failure		404	{object}	errcode.Error	"找不到页面"
//	@Failure		500	{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/tags/:id [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{Id: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%s", errors.Errors())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.ErrorF("app.DeleteTag errs:%v", errors)
		response.ToErrorResponseList(errcode.ErrorDeleteTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
