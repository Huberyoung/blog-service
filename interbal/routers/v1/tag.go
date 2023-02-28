package v1

import (
	"blog-service/global"
	"blog-service/interbal/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type Tag struct {
}

func NewTage() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {
	var tag model.Tag
	global.DBEngine.First(&tag)
	log.Printf("record %#v\n", tag)

	c.JSON(http.StatusOK, tag)
}

func (t Tag) List(c *gin.Context) {

}

func (t Tag) Create(c *gin.Context) {

}

func (t Tag) Update(c *gin.Context) {

}

func (t Tag) Delete(c *gin.Context) {

}
