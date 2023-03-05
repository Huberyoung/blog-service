package v1

import (
	"blog-service/global"
	"blog-service/interbal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/convert"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

// Get godoc
//
//	@Summary		获取单篇文章
//	@Description	通过唯一编号获取单篇文章
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			id	path		uint			true	"文章 ID"
//	@Success		200	{object}	model.Article	"成功"
//	@Failure		400	{object}	errcode.Error
//	@Failure		404	{object}	errcode.Error
//	@Failure		500	{object}	errcode.Error
//	@Router			/api/v1/articles/:id  [get]
func (a Article) Get(c *gin.Context) {
	param := service.GetArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errors)
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	article, err := svc.GetArticle(&param)
	if err != nil {
		global.Logger.ErrorF("svc.GetArticle err:%v", err)
		response.ToErrorResponseList(errcode.ErrorGetArticleFail)
		return
	}
	response.ToResponse(article)
}

// List godoc
//
//	@Summary		获取文章列表
//	@Description	通过文章标题，状态，以及分页情况获取文章列表
//	@Tags			Article
//	@Accept			json
//	@Produce		json
//	@Param			title		query		string			false	"文章标题"
//	@Param			state		query		uint			false	"文章状态 0 不可用，1可用"	Enums(0, 1)	default(1)
//	@Param			page		query		uint			false	"页码"
//	@Param			page_size	query		uint			false	"每页数量"
//	@Success		200			{object}	model.Article	"成功"
//	@Failure		400			{object}	errcode.Error
//	@Failure		404			{object}	errcode.Error
//	@Failure		500			{object}	errcode.Error
//	@Router			/api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ListArticleRequest{}

	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errors)
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}

	param2 := param
	totalRows, err := svc.CountArticle(&param2)
	if err != nil {
		global.Logger.ErrorF("svc.CountArticle err:%v", err)
		response.ToErrorResponseList(errcode.ErrorCountArticleListFail)
		return
	}

	articles, err := svc.ListArticle(&param, &pager)
	if err != nil {
		global.Logger.ErrorF("svc.ListArticle err:%v", err)
		response.ToErrorResponseList(errcode.ErrorGetArticleListFail)
		return
	}
	response.ToResponseList(articles, totalRows)
	return
}

// Create godoc
//
//	@Summary		新增文章
//	@Description	新增文章内容
//	@Tags			Article
//	@Produce		json
//	@Param			title			formData	string			true	"文章标题"		minlength(1)	maxlength(10)
//	@Param			desc			formData	string			true	"文章简述"		minlength(3)	maxlength(100)
//	@Param			content			formData	string			true	"文章内容"		minlength(3)	maxlength(10000)
//	@Param			cover_image_url	formData	string			true	"文章图片地址"	minlength(3)	maxlength(100)
//	@Param			state			formData	uint			false	"状态"		Enums(0, 1)		default(1)
//	@Param			created_by		formData	string			true	"创建者"		minlength(1)	maxlength(30)
//	@Success		200				{object}	model.Article	"成功"
//	@Failure		400				{object}	errcode.Error	"请求错误"
//	@Failure		404				{object}	errcode.Error	"找不到页面"
//	@Failure		500				{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}

	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%s", errors.Error())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateArticle(&param)
	if err != nil {
		global.Logger.ErrorF("app.CreateArticle errs:%v", err)
		response.ToErrorResponseList(errcode.ErrorCreateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Update godoc
//
//	@Summary		更新文章
//	@Description	更新文章内容
//	@Tags			Article
//	@Produce		json
//	@Param			title			formData	string			false	"文章标题"		minlength(1)	maxlength(10)
//	@Param			desc			formData	string			false	"文章简述"		minlength(3)	maxlength(100)
//	@Param			content			formData	string			false	"文章内容"		minlength(3)	maxlength(100)
//	@Param			cover_image_url	formData	string			false	"文章图片地址"	minlength(3)	maxlength(100)
//	@Param			state			formData	uint			false	"状态"		Enums(0, 1)		default(1)
//	@Param			modified_by		formData	string			true	"更新者"		minlength(1)	maxlength(30)
//	@Success		200				{object}	model.Article	"成功"
//	@Failure		400				{object}	errcode.Error	"请求错误"
//	@Failure		404				{object}	errcode.Error	"找不到页面"
//	@Failure		500				{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/articles [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt()}

	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%s", errors.Errors())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateArticle(&param)
	if err != nil {
		global.Logger.ErrorF("app.UpdateArticle errs:%v", errors)
		response.ToErrorResponseList(errcode.ErrorUpdateArticleFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

// Delete godoc
//
//	@Summary		删除文章
//	@Description	通过id删除文章
//	@Tags			Article
//	@Produce		json
//	@Param			id	path		uint			true	"文章ID"
//	@Success		200	{object}	model.Article	"成功"
//	@Failure		400	{object}	errcode.Error	"请求错误"
//	@Failure		404	{object}	errcode.Error	"找不到页面"
//	@Failure		500	{object}	errcode.Error	"内部错误"
//	@Router			/api/v1/articles/:id [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt()}
	response := app.NewResponse(c)
	valid, errors := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%s", errors.Errors())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(errors.Errors()...))
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteArticle(&param)
	if err != nil {
		global.Logger.ErrorF("app.DeleteArticle errs:%v", errors)
		response.ToErrorResponseList(errcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}
