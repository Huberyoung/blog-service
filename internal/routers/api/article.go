package api

import (
	"github.com/gin-gonic/gin"
	"service/global"
	"service/internal/service"
	"service/pkg/app"
	"service/pkg/convert"
	"service/pkg/errorcode"
)

type Article struct{}

func NewArticle() Article {
	return Article{}
}

// Get @Summary 获取单个文章
// @Produce json
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/api/articles [get]
func (a Article) Get(c *gin.Context) {
	param := service.ArticleGetRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	srv := service.New(c.Request.Context())
	article, err := srv.GetArticle(&param)
	if err != nil {
		global.Logger.ErrorF("svc.GetArticle err: %v", err)
		response.ToErrorResponse(errorcode.ErrorGetArticleListFail)
		return
	}
	response.ToResponse(article)
	return
}

// List @Summary 获取多个文章
// @Produce  json
// @Param title query string false "文章名称" maxlength(100)
// @Param state query int false "状态" Enums(0, 1) default(1)
// @Param page query int false "页码"
// @Param page_size query int false "每页数量"
// @Success 200 {object} model.ArticleSwagger "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/api/articles [get]
func (a Article) List(c *gin.Context) {
	param := service.ArticleListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	srv := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := srv.CountArticle(&service.CountArticleRequest{Title: param.Title, State: param.State})
	if err != nil {
		global.Logger.ErrorF("srv.CountArticle errs:%v", errs)
		response.ToErrorResponse(errorcode.ErrorCountArticleFail)
		return
	}

	Articles, err := srv.GetArticleList(&param, &pager)
	if err != nil {
		global.Logger.ErrorF("svc.GetArticleList err: %v", err)
		response.ToErrorResponse(errorcode.ErrorGetArticleListFail)
		return
	}

	response.ToResponseList(Articles, totalRows)
	return

}

// Create @Summary 新增文章
// @Produce  json
// @Param title body string true "文章名称" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(100) maxlength(65535)
// @Param state body int false "状态" Enums(0, 1) default(1)
// @Param created_by body string true "创建者" minlength(3) maxlength(100)
// @Success 200 {object} model.ArticleSwagger"成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/api/articles [post]
func (a Article) Create(c *gin.Context) {
	param := service.CreateArticleRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	srv := service.New(c.Request.Context())
	err := srv.CreateArticle(&param)
	if err != nil {
		global.Logger.ErrorF("srv.CreateArticle errs:%v", errs)
		response.ToErrorResponse(errorcode.ErrorCreateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// Update @Summary 更新文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Param title body string true "文章名称" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(100) maxlength(65535)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.ArticleSwagger"成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/api/articles/{id} [put]
func (a Article) Update(c *gin.Context) {
	param := service.UpdateArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	srv := service.New(c.Request.Context())
	err := srv.UpdateArticle(&param)
	if err != nil {
		global.Logger.ErrorF("srv.UpdateArticle errs:%v", errs)
		response.ToErrorResponse(errorcode.ErrorUpdateArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}

// Delete @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/api/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {
	param := service.DeleteArticleRequest{ID: convert.StrTo(c.Param("id")).MustUInt32()}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.ErrorF("app.BindAndValid errs:%v", errs)
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(errs.Errors()...))
		return
	}

	srv := service.New(c.Request.Context())
	err := srv.DeleteArticle(&param)
	if err != nil {
		global.Logger.ErrorF("srv.DeleteArticle errs:%v", errs)
		response.ToErrorResponse(errorcode.ErrorDeleteArticleFail)
		return
	}
	response.ToResponse(gin.H{})
}
