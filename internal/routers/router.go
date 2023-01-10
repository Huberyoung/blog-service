package routers

import (
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	_ "service/docs"
	"service/global"
	"service/internal/middleware"
	"service/internal/routers/api"
)

func NewRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.Translation(), middleware.JWT())
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	upload := NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	r.GET("/auth", api.GetAuth)

	article := api.NewArticle()
	tag := api.NewTag()

	group := r.Group("/api/api")
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
