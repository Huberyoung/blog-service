package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	_ "service/docs"
	v1 "service/internal/routers/v1"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery())
	url := ginSwagger.URL("http://127.0.0.1:8000/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	article := v1.NewArticle()
	tag := v1.NewTag()

	group := r.Group("/api/v1")
	{
		group.POST("/tags", tag.Create)
		group.DELETE("/tags/:id", tag.Delete)
		group.PUT("/tags/:id", tag.Update)
		group.PATCH("/tags/:id/state", tag.Update)
		group.GET("/tags", tag.List)

		group.POST("/articles", article.Create)
		group.DELETE("/articles/:id", article.Delete)
		group.PUT("/articles/:id", article.Update)
		group.PATCH("/articles/:id/state", article.Update)
		group.GET("/articles/:id", article.Get)
		group.GET("/articles", article.List)
	}
	return r
}
