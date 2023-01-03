package v1

import (
	"github.com/gin-gonic/gin"
	"service/pkg/app"
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
// @Router /api/v1/articles [get]
func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToResponse(gin.H{"ssdsd": "sfsgfsgs"})
	//app.NewResponse(c).ToErrorResponse(errorcode.UnauthorizedTokenError)
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
// @Router /api/v1/articles [get]
func (a Article) List(c *gin.Context) {
	//app.NewResponse(c).ToResponse(gin.H{"name": "这个是 List"})
	app.NewResponse(c).ToResponse(gin.H{"name": "这个是 List"})

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
// @Router /api/v1/articles [post]
func (a Article) Create(c *gin.Context) {}

// Update @Summary 更新文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Param title body string true "文章名称" minlength(3) maxlength(100)
// @Param content body string true "文章内容" minlength(100) maxlength(65535)
// @Param modified_by body string true "修改者" minlength(3) maxlength(100)
// @Success 200 {array} model.ArticleSwagger"成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles/{id} [put]
func (a Article) Update(c *gin.Context) {}

// Delete @Summary 删除文章
// @Produce  json
// @Param id path int true "文章 ID"
// @Success 200 {string} string "成功"
// @Failure 400 {object} errorcode.Error "请求错误"
// @Failure 500 {object} errorcode.Error "内部错误"
// @Router /api/v1/articles/{id} [delete]
func (a Article) Delete(c *gin.Context) {}
