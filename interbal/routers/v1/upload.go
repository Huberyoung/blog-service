package v1

import (
	"blog-service/global"
	"blog-service/interbal/service"
	"blog-service/pkg/app"
	"blog-service/pkg/convert"
	"blog-service/pkg/errcode"
	"blog-service/pkg/upload"
	"github.com/gin-gonic/gin"
)

type Upload struct {
}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(c *gin.Context) {
	response := app.NewResponse(c)

	file, header, err := c.Request.FormFile("file")
	if err != nil {
		global.Logger.ErrorF("Request.FormFile errs:%s", err.Error())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}

	fileType := convert.StrTo(c.PostForm("type")).MustUInt()
	if header == nil || fileType <= 0 {
		response.ToErrorResponseList(errcode.InvalidParams)
		return
	}

	svc := service.New(c.Request.Context())
	fileInfo, err := svc.UploadFile(upload.FileType(fileType), file, header)
	if err != nil {
		global.Logger.ErrorF("svc.UploadFile errs:%s", err.Error())
		response.ToErrorResponseList(errcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
