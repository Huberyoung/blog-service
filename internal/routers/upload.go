package routers

import (
	"github.com/gin-gonic/gin"
	"service/global"
	"service/internal/service"
	"service/pkg/app"
	"service/pkg/convert"
	"service/pkg/errorcode"
	"service/pkg/upload"
)

type Upload struct{}

func NewUpload() Upload {
	return Upload{}
}

func (u Upload) UploadFile(ctx *gin.Context) {
	response := app.NewResponse(ctx)
	file, header, err := ctx.Request.FormFile("file")
	fileType := convert.StrTo(ctx.PostForm("type")).MustInt()
	if err != nil {
		response.ToErrorResponse(errorcode.InvalidParams.WithDetails(err.Error()))
		return
	}
	if header == nil || fileType <= 0 {
		response.ToErrorResponse(errorcode.InvalidParams)
		return
	}
	s := service.New(ctx)
	fileInfo, err := s.UploadFile(upload.FileType(fileType), file, header)
	if err != nil {
		global.Logger.ErrorF("s.UploadFile err:%v", err)
		response.ToErrorResponse(errorcode.ErrorUploadFileFail.WithDetails(err.Error()))
	}

	response.ToResponse(gin.H{
		"file_access_url": fileInfo.AccessUrl,
	})
}
