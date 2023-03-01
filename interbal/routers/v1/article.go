package v1

import (
	"blog-service/pkg/app"
	"blog-service/pkg/errcode"
	"github.com/gin-gonic/gin"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	app.NewResponse(c).ToResponse(map[string]string{"title": "这里"})
	app.NewResponse(c).ToErrorResponseList(errcode.ServerError)
}

func (a Article) List(c *gin.Context) {

}

func (a Article) Create(c *gin.Context) {

}

func (a Article) Update(c *gin.Context) {

}

func (a Article) Delete(c *gin.Context) {

}
