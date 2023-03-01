package v1

import (
	"blog-service/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Article struct {
}

func NewArticle() Article {
	return Article{}
}

func (a Article) Get(c *gin.Context) {
	global.Logger.Info("这就是一个测试")
	global.Logger.InfoF("这就是[%d]个测试", 7)

	c.JSON(http.StatusOK, gin.H{"title": "测试"})
}

func (a Article) List(c *gin.Context) {

}

func (a Article) Create(c *gin.Context) {

}

func (a Article) Update(c *gin.Context) {

}

func (a Article) Delete(c *gin.Context) {

}
