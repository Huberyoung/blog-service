package routers

import (
	_ "blog-service/docs"
	"blog-service/global"
	"blog-service/interbal/middleware"
	v1 "blog-service/interbal/routers/v1"
	"blog-service/pkg/limiter"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.New()

	r.Use(middleware.AccessLog())
	r.Use(middleware.Recovery())
	methodLimiters := limiter.NewMethodLimiter().AddBuckets(limiter.LimiterBucketRule{
		Key:          "/user",
		FillInterval: time.Second,
		Capacity:     10,
		Quantum:      10,
	})
	r.Use(middleware.RateLimiter(methodLimiters))
	r.Use(middleware.ContextTimeout(global.AppSetting.DefaultContextTimeout))
	r.Use(middleware.Translations())
	r.Use(middleware.Tracing())
	
	user := v1.NewUser()
	r.POST("/user", user.Get)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
	tag := v1.NewTage()
	article := v1.NewArticle()

	upload := v1.NewUpload()
	r.POST("/upload/file", upload.UploadFile)
	r.StaticFS("/static", http.Dir(global.AppSetting.UploadSavePath))

	apiV1 := r.Group("/api/v1")
	apiV1.Use(middleware.JWT())
	{
		apiV1.POST("/tags", tag.Create)
		apiV1.DELETE("/tags/:id", tag.Delete)
		apiV1.PUT("/tags/:id", tag.Update)
		apiV1.PATCH("/tags/:id/state", tag.Update)
		apiV1.GET("/tags/:id", tag.Get)
		apiV1.GET("/tags", tag.List)

		apiV1.POST("/articles", article.Create)
		apiV1.DELETE("/articles/:id", article.Delete)
		apiV1.PUT("/articles/:id", article.Update)
		apiV1.PATCH("/articles/:id/state", article.Update)
		apiV1.GET("/articles/:id", article.Get)
		apiV1.GET("/articles", article.List)
	}

	return r
}
