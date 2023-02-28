package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Tag struct {
}

func NewTage() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"title": "这就是测试啊",
	})
}

func (t Tag) List(c *gin.Context) {

}

func (t Tag) Create(c *gin.Context) {

}

func (t Tag) Update(c *gin.Context) {

}

func (t Tag) Delete(c *gin.Context) {

}
